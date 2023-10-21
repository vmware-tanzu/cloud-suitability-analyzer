import pandas as pd

# Read data from CSV file
data = pd.read_csv('your_file.csv')

# Create pivot table
pivot_table = pd.pivot_table(data, values='value', index='application', columns='category', aggfunc='first')

# Fill missing values
pivot_table.fillna('NULL', inplace=True)

# Save pivot table to a new CSV file
pivot_table.to_csv('pivot_table.csv')

