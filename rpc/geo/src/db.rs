use std::env;
use tokio::sync::Mutex;
use lazy_static::lazy_static;
use deadpool_postgres::{Manager, Pool, tokio_postgres};
use serde::Deserialize;
use tokio_postgres::NoTls;
use config::{Config, File, FileFormat};

#[derive(Debug, Deserialize)]
struct Settings {
    database: DatabaseSettings,
}

#[derive(Debug, Deserialize)]
struct DatabaseSettings {
    host: String,
    user: String,
    password: String,
    dbname: String,
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
}

lazy_static! {
    pub static ref POOL: Mutex<Pool> = {
         // 从配置文件加载数据库连接信息
        let settings = Settings::new();
        let db = settings.database;
        // 创建连接池管理器
        let mut cfg = tokio_postgres::Config::new();
        cfg.host(&db.host);
        cfg.user(&db.user);
        cfg.password(&db.password);
        cfg.dbname(&db.dbname);

        let manager = Manager::new(cfg, NoTls);
        let pool = Pool::builder(manager)
            .max_size(10)
            .build()
            .unwrap();

        Mutex::new(pool)
    };
}