SELECT DISTINCT
       application,
       category,
       value
  FROM findings
  WHERE category != "File Finding" and
        category != "SLOC"

