use std::env;
use crate::programs::config::CPLConfigLoader;
use anyhow::Result;



fn init_gdal_config() -> Result<()> {
    let config_path = env::current_dir()
        .map(|path| path.join("config/gdal_config.ini"))
        .unwrap_or_else(|_| "config/gdal_config.ini".into());
    CPLConfigLoader::load_config_options_from_file(config_path, Some(true))
        .map_err(|e| anyhow::anyhow!("Failed to load GDAL config: {}", e))?;
    Ok(())
}