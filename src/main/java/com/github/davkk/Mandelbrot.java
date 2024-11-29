package com.github.davkk;

import java.awt.Color;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

public class Mandelbrot {
    public static BufferedImage mandelbrot(
            int size,
            double minRe,
            double maxRe,
            double minIm,
            double maxIm,
            int maxIterations) {

        BufferedImage image = new BufferedImage(size, size, BufferedImage.TYPE_INT_RGB);

        double pxRe = (maxRe - minRe) / (size - 1);
        double pxIm = (maxIm - minIm) / (size - 1);

        for (int y = 0; y < size; y++) {
            double cIm = maxIm - y * pxIm;

            for (int x = 0; x < size; x++) {
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

        return image;
    }

    public static void benchmark(int size, int repeat) {
        var start = System.nanoTime();
        for (int i = 0; i < repeat; i++) {
            mandelbrot(size, -2.1, 0.6, -1.2, 1.2, 200);
        }
        var end = System.nanoTime();
        var time = (end - start) / 1e9;
        System.out.printf("%d %f\n", size, time / (float) repeat);
    }

    public static void main(String[] args) throws IOException {
        assert args.length > 0;

        var size = Integer.parseInt(args[0]);

        benchmark(size, 10);

        BufferedImage image = mandelbrot(size, -0.7435, -0.7395, 0.1312, 0.1352, 200);
        ImageIO.write(image, "png", new File(String.format("output/mandelbrot_%d.png", size)));
    }
}
