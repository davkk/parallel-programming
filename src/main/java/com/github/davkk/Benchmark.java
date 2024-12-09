package com.github.davkk;

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

import javax.imageio.ImageIO;

public class Benchmark {
    final public static double minRe = -2.1;
    final public static double maxRe = 0.6;
    final public static double minIm = -1.2;
    final public static double maxIm = 1.2;
    final public static int maxIterations = 200;

    final public static int cores = Runtime.getRuntime().availableProcessors();

    public static void main(String[] args) throws InterruptedException {
        assert args.length > 0;

        var type = args[0];
        var size = Integer.parseInt(args[1]);

        var pool = Executors.newFixedThreadPool(cores);

        Consumer<BufferedImage> runFunc = switch (type) {
            case "single" -> Benchmark::runSingle;
            case "threads" -> Benchmark::runThreads;
            case "pool-once" -> (image) -> runPool(image, pool);
            case "pool-every" -> Benchmark::runPool;
            case "pool-blocks-less" -> (image) -> runPool(image, cores / 2);
            case "pool-blocks-more" -> (image) -> runPool(image, cores * 2);
            default -> throw new IllegalArgumentException(type);
        };

        var image = new BufferedImage(size, size, BufferedImage.TYPE_INT_RGB);
        var repeat = 20;

        var start = System.nanoTime();
        for (int i = 0; i < repeat; i++) {
            runFunc.accept(image);
        }
        var end = System.nanoTime();
        var time = (end - start) / 1e9;

        System.out.printf("%d %f\n", size, time / (float) repeat);

        pool.shutdown();

        try {
            if (!pool.awaitTermination(1, TimeUnit.MINUTES)) {
                pool.shutdownNow();
            }
        } catch (InterruptedException e) {
            pool.shutdownNow();
            Thread.currentThread().interrupt();
        }

        try {
            ImageIO.write(image, "png", new File("test.png"));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void runSingle(BufferedImage image) {
        Mandelbrot.generate(image, 0, image.getWidth(), minRe, maxRe, minIm, maxIm, maxIterations);
    }

    private static void runThreads(BufferedImage image) {
        var size = image.getWidth();
        var blockSize = size / cores;

        var threads = new Thread[cores];

        for (var y = 0; y < cores; y++) {
            var y0 = y * blockSize;
            threads[y] = new Thread(() -> {
                Mandelbrot.generate(image, y0, blockSize, minRe, maxRe, minIm, maxIm, 200);
            });
        }

        for (var thread : threads) {
            thread.start();
        }

        for (var thread : threads) {
            try {
                thread.join();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                throw new RuntimeException(e);
            }
        }
    }

    private static void runPool(BufferedImage image) {
        var size = image.getWidth();
        var blockSize = size / cores;

        var pool = Executors.newFixedThreadPool(cores);

        for (var y = 0; y < cores; y++) {
            var y0 = y * blockSize;
            pool.execute(() -> {
                Mandelbrot.generate(image, y0, blockSize, minRe, maxRe, minIm, maxIm, 200);
            });
        }

        pool.shutdown();

        try {
            if (!pool.awaitTermination(1, TimeUnit.MINUTES)) {
                pool.shutdownNow();
            }
        } catch (InterruptedException e) {
            pool.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }

    private static void runPool(BufferedImage image, ExecutorService pool) {
        var size = image.getWidth();
        var blockSize = size / cores;

        for (var y = 0; y < cores; y++) {
            var y0 = y * blockSize;
            pool.execute(() -> {
                Mandelbrot.generate(image, y0, blockSize, minRe, maxRe, minIm, maxIm, 200);
            });
        }
    }

    private static void runPool(BufferedImage image, int blocks) {
        var size = image.getWidth();
        var blockSize = size / blocks;

        var pool = Executors.newFixedThreadPool(cores);

        for (var y = 0; y < blocks; y++) {
            var y0 = y * blockSize;
            pool.execute(() -> {
                Mandelbrot.generate(image, y0, blockSize, minRe, maxRe, minIm, maxIm, 200);
            });
        }

        pool.shutdown();

        try {
            if (!pool.awaitTermination(1, TimeUnit.MINUTES)) {
                pool.shutdownNow();
            }
        } catch (InterruptedException e) {
            pool.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }
}
