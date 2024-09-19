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
        let settings = Settings::new();
        let db = settings.database;

        let cfg = tokio_postgres::Config::new()
            .host(&db.host)
            .user(&db.user)
            .password(&db.password)
            .dbname(&db.dbname);

        let manager = Manager::new(cfg, NoTls);
        let pool = Pool::builder(manager)
            .max_size(10)
            .build()
            .expect("连接池创建失败");

        Mutex::new(pool)
    };
}