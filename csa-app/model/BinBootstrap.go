/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapBinsTemplate.txt found under go/resources folder
//Created @ 2024-03-13 11:06:12.440394 -0500 CDT m=+0.686364906

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
            []*BinTag{  { Name: "modernize", Type: 1, Action: "OR", },  },
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
            []*BinTag{  { Name: "cache", Type: 1, Action: "OR", },  },
             },
        
            { Name: "FAAS",
            Tags:
            []*BinTag{  { Name: "faas", Type: 1, Action: "OR", },  },
             },
        
            { Name: "DATABASE",
            Tags:
            []*BinTag{  { Name: "database", Type: 1, Action: "OR", },  },
             },
        
            { Name: "MESSAGING",
            Tags:
            []*BinTag{  { Name: "messaging", Type: 1, Action: "OR", },  },
             },
        
            { Name: "SECURITY",
            Tags:
            []*BinTag{  { Name: "security", Type: 1, Action: "OR", },  },
             },
        
            { Name: "LOGGING",
            Tags:
            []*BinTag{  { Name: "logging", Type: 1, Action: "OR", },  },
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
        
    }
    return BootstrapBins
}