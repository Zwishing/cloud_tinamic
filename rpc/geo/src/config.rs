use std::{env, sync::OnceLock};
use deadpool_postgres::{Manager, Pool, tokio_postgres};
use serde::Deserialize;
use tokio::sync::Mutex;
use tokio_postgres::NoTls;
use config::{Config, File, FileFormat};
use minio::s3::client::Client as MinioClient;
use minio::s3::ClientBuilder;
use minio::s3::creds::StaticProvider;

#[derive(Debug,Clone,Deserialize)]
pub struct Settings {
    database: DatabaseSettings,
    minio: MinioSettings
}

#[derive(Debug, Clone,Deserialize)]
struct DatabaseSettings {
    host: String,
    user: String,
    password: String,
    dbname: String,
    port:u16,
}

#[derive(Debug, Clone, Deserialize)]
struct MinioSettings {
    endpoint:String,
    access_key:String,
    secret_key:String,
}

impl Settings {
    fn new() -> Self {
        let config_path = env::current_dir()
            .map(|path| path.join("config/config.toml"))
            .unwrap_or_else(|_| "config/config.toml".into());

        Config::builder()
            .add_source(File::new(config_path.to_str().unwrap(), FileFormat::Toml))
            .build()
            .expect("配置文件加载失败")
            .try_deserialize()
            .expect("配置反序列化失败")
    }

    // 将配置转换为 tokio_postgres 配置
    fn to_postgres_config(&self) -> tokio_postgres::Config {
        let mut cfg = tokio_postgres::Config::new();
        let db = &self.database;
        cfg.host(&db.host);
        cfg.user(&db.user);
        cfg.password(&db.password);
        cfg.dbname(&db.dbname);
        cfg.port(db.port);
        cfg
    }
    // to_connection_string(): string{
    pub fn to_connection_string(&self) -> String {  
        format!(
            "postgresql://{}:{}@{}:{}/{}",
            self.database.user,
            self.database.password,
            self.database.host,
            self.database.port,
            self.database.dbname
        )
    }

    // PG: host=1.92.113.25 user=postgres password=admin321 dbname=tinamic
    pub fn to_pg_string(&self) -> String {
        format!(
            "PG: host={} user={} password={} dbname={} port={}",
            self.database.host,
            self.database.user,
            self.database.password,
            self.database.dbname,
            self.database.port,
        )
    }
}


// 使用 OnceLock 进行同步延迟初始化
static DB_SETTINGS: OnceLock<Mutex<Settings>> = OnceLock::new();
static POOL: OnceLock<Mutex<Pool>> = OnceLock::new();

static MINIO_CLIENT: OnceLock<Mutex<MinioClient>> = OnceLock::new();

pub fn init_settings() -> &'static Mutex<Settings> {
    DB_SETTINGS.get_or_init(|| Mutex::new(Settings::new()))
}

fn init_pool() -> &'static Mutex<Pool> {
    POOL.get_or_init(|| {
        let settings = init_settings().blocking_lock().clone();
        let config = settings.to_postgres_config();

        // 创建连接池管理器
        let manager = Manager::new(config, NoTls);
        let pool = Pool::builder(manager)
            .max_size(20)
            .build()
            .unwrap_or_else(|e| {
                tracing::error!("Failed to build database pool: {}", e);
                panic!("Database pool initialization failed")
            });

        tracing::info!("Successfully connected to PostgreSQL database @ {}", settings.database.host);
        Mutex::new(pool)
    })
}

fn init_minio_client() -> &'static Mutex<MinioClient> {
    MINIO_CLIENT.get_or_init(|| {
        let settings = init_settings().blocking_lock().clone();
        let static_provider = StaticProvider::new(
            &settings.minio.access_key,
            &settings.minio.secret_key,
            None,
        );

        let client = ClientBuilder::new(
            settings.minio.endpoint.parse().expect("Invalid MinIO endpoint URL")
        ).provider(Some(Box::new(static_provider)))
            .build()
            .unwrap_or_else(|e| {
                tracing::error!("Failed to build MinIO client: {}", e);
                panic!("MinIO client initialization failed")
            });
        tracing::info!("Successfully connected to MinIO @ {}", settings.minio.endpoint);
        Mutex::new(client)
    })
}

// 获取连接池的公共函数
pub async fn get_pool() -> &'static Mutex<Pool> {
    init_pool()
}

// 获取数据库设置的公共函数
pub fn get_settings() -> &'static Mutex<Settings> {
    init_settings()
}

// 获取MinIO客户端的公共函数
pub fn get_minio_client() -> &'static Mutex<MinioClient> {
    init_minio_client()
}