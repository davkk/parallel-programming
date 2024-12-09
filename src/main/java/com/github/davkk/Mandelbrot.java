package com.github.davkk;

import java.awt.Color;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

public class Mandelbrot {
    public static void generate(
            BufferedImage image,
            int y0,
            double size,
            double minRe,
            double maxRe,
            double minIm,
            double maxIm,
            int maxIterations //
    ) {
        double pxRe = (maxRe - minRe) / (image.getWidth() - 1);
        double pxIm = (maxIm - minIm) / (image.getHeight() - 1);

        for (int y = y0; y < Math.min(y0 + size, image.getHeight()); y++) {
            double cIm = maxIm - y * pxIm;

            for (int x = 0; x < image.getWidth(); x++) {
                double cRe = minRe + x * pxRe;

                double zRe = 0;
                double zIm = 0;
                int iterations = 0;

                while (zRe * zRe + zIm * zIm <= 4 && iterations < maxIterations) {
                    double nextRe = zRe * zRe - zIm * zIm + cRe;
                    double nextIm = 2 * zRe * zIm + cIm;

                    zRe = nextRe;
                    zIm = nextIm;

                    iterations++;
                }

                float hue = iterations / (float) maxIterations * 2.5f;
                float saturation = 1.0f;
                float brightness = iterations < maxIterations ? 1.0f : 0.0f;

                int color = Color.HSBtoRGB(hue, saturation, brightness);
                image.setRGB(x, y, color);
            }
        }
    }

    public static void main(String[] args) throws IOException {
        assert args.length > 0;
        var size = Integer.parseInt(args[0]);

        BufferedImage image = new BufferedImage(size, size, BufferedImage.TYPE_INT_RGB);
        generate(image, 0, size, -0.7435, -0.7395, 0.1312, 0.1352, 200);

        ImageIO.write(image, "png", new File(String.format("output/mandelbrot_%d.png", size)));
    }
}
