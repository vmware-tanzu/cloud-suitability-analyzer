/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapBinsTemplate.txt found under go/resources folder
//Created @ 2023-09-22 15:55:43.964827 -0500 CDT m=+0.361750980

func BootstrapBins() []Bin {
    var BootstrapBins = []Bin{
        
            { Name: "TKG",
            Tags:
            []*BinTag{  { Name: "docker", Type: 1, Action: "OR", },  { Name: "stateful", Type: 0, Action: "AND", },  { Name: "javaee", Type: 0, Action: "AND", },  { Name: "full-profile", Type: 0, Action: "AND", },  { Name: "jni", Type: 1, Action: "OR", },  { Name: "non-standard-protocol", Type: 1, Action: "OR", },  { Name: "corba", Type: 1, Action: "OR", },  },
             },
        
            { Name: "TAS",
            Tags:
            []*BinTag{  { Name: "web-profile", Type: 1, Action: "OR", },  { Name: "spring", Type: 1, Action: "OR", },  { Name: "spring-boot", Type: 1, Action: "OR", },  { Name: "web-container", Type: 1, Action: "OR", },  { Name: "rest", Type: 1, Action: "OR", },  },
             },
        
            { Name: "BATCH",
            Tags:
            []*BinTag{  { Name: "batch", Type: 1, Action: "OR", },  { Name: "etl", Type: 1, Action: "OR", },  },
             },
        
            { Name: "CONTAINER",
            Tags:
            []*BinTag{  { Name: "term", Type: 1, Action: "OR", },  { Name: "metrics", Type: 1, Action: "OR", },  { Name: "sudo", Type: 1, Action: "OR", },  { Name: "health-check", Type: 1, Action: "OR", },  { Name: "ehcache", Type: 1, Action: "OR", },  { Name: "non-root-user", Type: 1, Action: "OR", },  { Name: "hard-ip", Type: 1, Action: "OR", },  { Name: "process-exit", Type: 1, Action: "OR", },  { Name: "dist-cache", Type: 1, Action: "OR", },  { Name: "ws-cluster", Type: 1, Action: "OR", },  { Name: "wl-cluster", Type: 1, Action: "OR", },  { Name: "io", Type: 1, Action: "OR", },  { Name: "log2file", Type: 1, Action: "OR", },  { Name: "docker", Type: 1, Action: "OR", },  { Name: "transaction", Type: 1, Action: "OR", },  },
             },
        
            { Name: "BOOT",
            Tags:
            []*BinTag{  { Name: "cdi", Type: 1, Action: "OR", },  { Name: "ejb", Type: 1, Action: "OR", },  { Name: "jax-ws", Type: 1, Action: "OR", },  { Name: "jsf", Type: 1, Action: "OR", },  { Name: "mdb", Type: 1, Action: "OR", },  { Name: "struts", Type: 1, Action: "OR", },  { Name: "txn", Type: 1, Action: "OR", },  { Name: "servlet", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MICRO-SERVICE",
            Tags:
            []*BinTag{  { Name: "spring-boot", Type: 0, Action: "AND", },  },
             },
        
            { Name: "SPRING",
            Tags:
            []*BinTag{  { Name: "spring", Type: 1, Action: "OR", },  },
             },
        
            { Name: "STATEFUL",
            Tags:
            []*BinTag{  { Name: "stateful", Type: 1, Action: "OR", },  { Name: "spring", Type: 2, Action: "EXCLUDE", },  },
             },
        
            { Name: "WEB-MVC",
            Tags:
            []*BinTag{  { Name: "ui", Type: 1, Action: "OR", },  { Name: "web-app", Type: 0, Action: "AND", },  { Name: "web-container", Type: 0, Action: "AND", },  { Name: "jsp", Type: 1, Action: "OR", },  { Name: "jsf", Type: 1, Action: "OR", },  { Name: "portlet", Type: 1, Action: "OR", },  { Name: "struts", Type: 1, Action: "OR", },  { Name: "jetty", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DESKTOP-APP",
            Tags:
            []*BinTag{  { Name: "desktop-app", Type: 1, Action: "OR", },  },
             },
        
            { Name: "CACHING",
            Tags:
            []*BinTag{  { Name: "cache", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MONOLITH",
            Tags:
            []*BinTag{  { Name: "soap", Type: 0, Action: "AND", },  { Name: "web-service", Type: 0, Action: "AND", },  { Name: "messaging", Type: 0, Action: "AND", },  { Name: "ear", Type: 0, Action: "AND", },  { Name: "stateful", Type: 0, Action: "AND", },  { Name: "rpc", Type: 1, Action: "OR", },  { Name: "ant", Type: 1, Action: "OR", },  { Name: "websphere", Type: 1, Action: "OR", },  { Name: "weblogic", Type: 1, Action: "OR", },  { Name: "glassfish", Type: 1, Action: "OR", },  { Name: "jboss", Type: 1, Action: "OR", },  { Name: "security", Type: 1, Action: "OR", },  },
             },
        
            { Name: "FAAS",
            Tags:
            []*BinTag{  { Name: "io", Type: 1, Action: "OR", },  { Name: "security", Type: 1, Action: "OR", },  { Name: "faas", Type: 1, Action: "OR", },  },
             },
        
    }
    return BootstrapBins
}