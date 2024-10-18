use std::ffi::CString;
use std::mem::ManuallyDrop;
use std::path::{Path, PathBuf};
use gdal::Dataset;
use gdal::errors::{GdalError,Result};
use super::util::_path_to_c_string;


/// Path or Dataset to store Vectors or Raster
pub enum DatasetDestination {
    Path(CString),
    Dataset {
        dataset: ManuallyDrop<Dataset>,
        drop: bool,
    },
}

impl TryFrom<&str> for DatasetDestination {
    type Error = GdalError;

    fn try_from(path: &str) ->Result<Self> {
        Self::path(path)
    }
}

impl TryFrom<&Path> for DatasetDestination {
    type Error = GdalError;

    fn try_from(path: &Path) -> Result<Self> {
        Self::path(path)
    }
}

impl TryFrom<PathBuf> for DatasetDestination {
    type Error = GdalError;

    fn try_from(path: PathBuf) -> Result<Self> {
        Self::path(path)
    }
}

impl From<Dataset> for DatasetDestination {
    fn from(dataset: Dataset) -> Self {
        Self::dataset(dataset)
    }
}

impl Drop for DatasetDestination {
    fn drop(&mut self) {
        match self {
            Self::Path(_) => {}
            Self::Dataset { dataset, drop } => {
                if *drop {
                    unsafe {
                        ManuallyDrop::drop(dataset);
                    }
                }
            }
        }
    }
}

impl DatasetDestination {
    pub fn dataset(dataset: Dataset) -> Self {
        Self::Dataset {
            dataset: ManuallyDrop::new(dataset),
            drop: true,
        }
    }

    pub fn path<P: AsRef<Path>>(path: P) -> Result<Self> {
        let c_path = _path_to_c_string(path.as_ref())?;
        Ok(Self::Path(c_path))
    }

    pub unsafe fn do_no_drop_dataset(&mut self) {
        match self {
            Self::Path(_) => {}
            Self::Dataset { dataset: _, drop } => {
                *drop = false;
            }
        }
    }
}

