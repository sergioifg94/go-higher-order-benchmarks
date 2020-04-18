library(ggplot2)
library(readr)

bench_results <- read_csv("result.csv")

plot_nsperop <- ggplot(bench_results, aes(size, nsperop, colour = name)) +
    geom_point() +
    geom_line() +
    scale_x_log10() +
    facet_grid(rows = vars(ncalls))

ggsave("plot_nsperop.png", width = 7, height = 14)

plot_bperop <- ggplot(bench_results, aes(size, bperop, colour = name)) +
    geom_point() +
    geom_line() +
    scale_x_log10() +
    facet_grid(rows = vars(ncalls))

ggsave("plot_bperop.png", width = 7, height = 14)