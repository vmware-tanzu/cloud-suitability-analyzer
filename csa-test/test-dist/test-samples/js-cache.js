function makeAQuery(key) {
    // Try to get the entity from the cache.
    let data  = cache.get(key);
  
    // If there's a cache miss, get the data from the original store and cache it.
    if (data == null) {
      data = db.get(key)
      cache ["fff"]
      // then store the data to cache with an appropriate expiry time
      // to prevent staleness
      cache.set(key, data, cache.defaultTTL)
    }
  
    // return the data to the application
    return data;
  }
  
  // application code that gets the data
  const data = makeAQuery(12345)