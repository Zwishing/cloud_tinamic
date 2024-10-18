use std::env;
use crate::programs::config::CPLConfigLoader;
use anyhow::Result;



pub fn init_gdal_config() -> Result<()> {
    let config_path = env::current_dir()
        .map(|path| path.join("config/gdal_config.ini"))
        .unwrap_or_else(|_| "config/gdal_config.ini".into());
    CPLConfigLoader::load_config_options_from_file(config_path, Some(true))
        .map_err(|e| anyhow::anyhow!("Failed to load GDAL config: {}", e))?;
    Ok(())
}

#[cfg(test)]
mod tsets{
    use crate::gdal_config::init_gdal_config;
    use std::env;
    use std::path::PathBuf;

    #[test]
    fn test_init_gdal_config() {

        // Run the function and check the result
        let result = init_gdal_config();
        assert!(result.is_ok(), "init_gdal_config failed: {:?}", result.err());
    }
}