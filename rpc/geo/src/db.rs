use std::env;
use tokio::sync::Mutex;
use lazy_static::lazy_static;
use deadpool_postgres::{Manager, Pool, tokio_postgres};
use serde::Deserialize;
use tokio_postgres::{NoTls};
use config::{Config, File, FileFormat};


// 定义配置结构体，用于存储从配置文件中读取的数据
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
        // 设置配置文件路径，可以通过环境变量或命令行参数传递
        let config_path = format!("{}/config/config.toml", env::current_dir().unwrap().display());
        let builder = Config::builder()
            .add_source(File::new(config_path.as_str(),FileFormat::Toml));
        builder.build().unwrap().try_deserialize().unwrap()
    }
}

lazy_static!{
    pub static ref POOL:Mutex<Pool> = {

        // 从配置文件加载数据库连接信息
        let settings = Settings::new();
        let database_settings = settings.database;
        // 创建连接池管理器
        let mut cfg = tokio_postgres::Config::new();
        cfg.host(&database_settings.host);
        cfg.user(&database_settings.user);
        cfg.password(&database_settings.password);
        cfg.dbname(&database_settings.dbname);

        let manager = Manager::new(cfg, NoTls);
        let pool = Pool::builder(manager)
            .max_size(10)
            .build()
            .unwrap();

        Mutex::new(pool)
    };
}