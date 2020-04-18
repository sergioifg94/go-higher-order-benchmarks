# Run the benchmark
echo "Running benchmark..."
go run main.go > result.csv

# Plot the results
echo "Plotting results..."
Rscript plot_results.R