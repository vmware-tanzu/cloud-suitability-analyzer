/*
 * Licensed to the Apache Software Foundation (ASF) under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for additional information regarding
 * copyright ownership. The ASF licenses this file to You under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance with the License. You may obtain a
 * copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */
package org.apache.geode_examples.listener;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;
import java.util.Map;
import java.util.Queue;
import java.util.Random;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.stream.IntStream;

import org.apache.geode.cache.CacheListener;
import org.apache.geode.cache.EntryEvent;
import org.apache.geode.cache.Region;
import org.apache.geode.cache.client.ClientCache;
import org.apache.geode.cache.client.ClientCacheFactory;
import org.apache.geode.cache.client.ClientRegionFactory;
import org.apache.geode.cache.client.ClientRegionShortcut;

public class Example {
  public static final int ITERATIONS = 100;

  public static void main(String[] args) {
    // connect to the locator using default port 10334
    ClientCache cache = new ClientCacheFactory().addPoolLocator("127.0.0.1", 10334)
        .set("log-level", "WARN").create();

    Example example = new Example();

    // create a local region that matches the server region
    ClientRegionFactory<Integer, String> clientRegionFactory =
        cache.createClientRegionFactory(ClientRegionShortcut.PROXY);
    clientRegionFactory.addCacheListener(new ExampleCacheListener());
    Region<Integer, String> region = clientRegionFactory.create("example-region");

    example.putEntries(region);
    cache.close();
  }

  private Collection<Integer> generateIntegers() {
    IntStream stream = new Random().ints(0, ITERATIONS);
    Iterator<Integer> iterator = stream.iterator();
    Collection<Integer> integers = new ArrayList<>();
    while (iterator.hasNext() && integers.size() < ITERATIONS) {
      Integer integer = iterator.next();
      if (!integers.contains(integer)) {
        integers.add(integer);
      }
    }
    return integers;
  }

  public void putEntries(Map<Integer, String> region) {
    Collection<Integer> integers = generateIntegers();
    Iterator<Integer> iterator = integers.iterator();

    while (iterator.hasNext()) {
      Integer integer = iterator.next();
      region.put(integer, integer.toString());
    }
    System.out.println("Created " + integers.size() + " entries.");
  }
}