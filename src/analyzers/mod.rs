mod hash;

use types::AnalyzedImage;
use magick_rust::MagickWand;

pub fn run(source: &Vec<u8>) -> Result<AnalyzedImage, &'static str> {
    let wand = MagickWand::new();
    wand.read_image_blob(source)?;

    let format = wand.get_image_format()?;
    let (width, height) = get_dimensions(wand, &format);
    let hash = hash::perceptual_hash(source)?;

    Ok(AnalyzedImage {
        width: width,
        height: height,
        hash: format!("{:064b}", hash)
    })
}

fn get_dimensions(wand: MagickWand, format: &str) -> (usize, usize) {
    match format {
        "GIF" => {
            let (width, height, _, _) = wand.get_image_page();
            (width, height)
        }
        _ => {
            (wand.get_image_width(), wand.get_image_height())
        }
    }
}
