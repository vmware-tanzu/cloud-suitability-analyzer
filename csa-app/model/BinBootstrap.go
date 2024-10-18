/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapBinsTemplate.txt found under go/resources folder
//Created @ 2024-10-18 10:32:29.65452 -0500 CDT m=+0.604860137

func BootstrapBins() []Bin {
    var BootstrapBins = []Bin{
        
            { Name: "",
            Tags:
            []*BinTag{  },
             },
        
            { Name: "TAS",
            Tags:
            []*BinTag{  { Name: "tas", Type: 1, Action: "OR", },  { Name: "env-config", Type: 1, Action: "OR", },  { Name: "dependencies", Type: 1, Action: "OR", },  { Name: "spring", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MODERNIZE",
            Tags:
            []*BinTag{  { Name: "modernize", Type: 1, Action: "OR", },  { Name: "weblogic", Type: 1, Action: "OR", },  { Name: "glassfish", Type: 1, Action: "OR", },  { Name: "jboss", Type: 1, Action: "OR", },  { Name: "websphere", Type: 1, Action: "OR", },  },
             },
        
            { Name: "BATCH",
            Tags:
            []*BinTag{  { Name: "batch", Type: 1, Action: "OR", },  },
             },
        
            { Name: "CONTAINER",
            Tags:
            []*BinTag{  { Name: "container", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DEPENDENCIES",
            Tags:
            []*BinTag{  { Name: "dependencies", Type: 1, Action: "OR", },  },
             },
        
            { Name: "SPRING",
            Tags:
            []*BinTag{  { Name: "spring", Type: 1, Action: "OR", },  { Name: "spring-boot", Type: 1, Action: "OR", },  },
             },
        
            { Name: "STATEFUL",
            Tags:
            []*BinTag{  { Name: "stateful", Type: 1, Action: "OR", },  { Name: "spring", Type: 2, Action: "EXCLUDE", },  },
             },
        
            { Name: "FRONTEND",
            Tags:
            []*BinTag{  { Name: "frontend", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DESKTOP-APP",
            Tags:
            []*BinTag{  { Name: "desktop-app", Type: 1, Action: "OR", },  },
             },
        
            { Name: "CACHING",
            Tags:
            []*BinTag{  { Name: "cache", Type: 1, Action: "OR", },  { Name: "memcache", Type: 1, Action: "OR", },  { Name: "hazelcast", Type: 1, Action: "OR", },  { Name: "apache-geode", Type: 1, Action: "OR", },  { Name: "spring-boot-cache", Type: 1, Action: "OR", },  { Name: "redis", Type: 1, Action: "OR", },  { Name: "ehcache", Type: 1, Action: "OR", },  },
             },
        
            { Name: "FAAS",
            Tags:
            []*BinTag{  { Name: "faas", Type: 1, Action: "OR", },  },
             },
        
            { Name: "WEBSERVER",
            Tags:
            []*BinTag{  { Name: "webserver", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DATABASE",
            Tags:
            []*BinTag{  { Name: "database", Type: 1, Action: "OR", },  { Name: "ibm-db2", Type: 1, Action: "OR", },  { Name: "elasticsearch", Type: 1, Action: "OR", },  { Name: "graphql", Type: 1, Action: "OR", },  { Name: "mongodb", Type: 1, Action: "OR", },  { Name: "mssql", Type: 1, Action: "OR", },  { Name: "oracle", Type: 1, Action: "OR", },  { Name: "cassandra", Type: 1, Action: "OR", },  { Name: "hadoop", Type: 1, Action: "OR", },  { Name: "mysql", Type: 1, Action: "OR", },  { Name: "postgres", Type: 1, Action: "OR", },  { Name: "sqlite", Type: 1, Action: "OR", },  { Name: "peewee", Type: 1, Action: "OR", },  { Name: "ssis", Type: 1, Action: "OR", },  { Name: "sqlserver", Type: 1, Action: "OR", },  { Name: "mariadb", Type: 1, Action: "OR", },  { Name: "apache-ignite", Type: 1, Action: "OR", },  { Name: "couchbase", Type: 1, Action: "OR", },  { Name: "elasticsearch", Type: 1, Action: "OR", },  { Name: "rdbms", Type: 1, Action: "OR", },  { Name: "nosql", Type: 1, Action: "OR", },  { Name: "in-memory", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MESSAGING",
            Tags:
            []*BinTag{  { Name: "messaging", Type: 1, Action: "OR", },  { Name: "activemq", Type: 1, Action: "OR", },  { Name: "apache-flume", Type: 1, Action: "OR", },  { Name: "apache-storm", Type: 1, Action: "OR", },  { Name: "ibm-mq", Type: 1, Action: "OR", },  { Name: "msmq", Type: 1, Action: "OR", },  { Name: "kafka", Type: 1, Action: "OR", },  { Name: "solace", Type: 1, Action: "OR", },  { Name: "masstransit", Type: 1, Action: "OR", },  { Name: "nservicebus", Type: 1, Action: "OR", },  { Name: "apache-pulsar", Type: 1, Action: "OR", },  { Name: "rocketmq", Type: 1, Action: "OR", },  { Name: "zeromq", Type: 1, Action: "OR", },  { Name: "nats", Type: 1, Action: "OR", },  { Name: "rabbitmq", Type: 1, Action: "OR", },  { Name: "apache-flink", Type: 1, Action: "OR", },  { Name: "jms", Type: 1, Action: "OR", },  { Name: "webspheremq", Type: 1, Action: "OR", },  { Name: "webspheremq", Type: 1, Action: "OR", },  { Name: "webspheremq", Type: 1, Action: "OR", },  { Name: "webspheremq", Type: 1, Action: "OR", },  },
             },
        
            { Name: "SECURITY",
            Tags:
            []*BinTag{  { Name: "security", Type: 1, Action: "OR", },  },
             },
        
            { Name: "LOGGING",
            Tags:
            []*BinTag{  { Name: "logging", Type: 1, Action: "OR", },  },
             },
        
            { Name: "SCHEDULING",
            Tags:
            []*BinTag{  { Name: "scheduler", Type: 1, Action: "OR", },  },
             },
        
            { Name: "CONFIGURATION",
            Tags:
            []*BinTag{  { Name: "configuration", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MONITORING",
            Tags:
            []*BinTag{  { Name: "monitoring", Type: 1, Action: "OR", },  },
             },
        
            { Name: "BUILD-SYSTEM",
            Tags:
            []*BinTag{  { Name: "build-system", Type: 1, Action: "OR", },  },
             },
        
            { Name: "VERSION",
            Tags:
            []*BinTag{  { Name: "java-version", Type: 1, Action: "OR", },  { Name: "spring-boot-version", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DOCUMENTATION",
            Tags:
            []*BinTag{  { Name: "documentation", Type: 1, Action: "OR", },  },
             },
        
    }
    return BootstrapBins
}