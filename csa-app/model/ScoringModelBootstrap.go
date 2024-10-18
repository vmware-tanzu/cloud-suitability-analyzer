/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapScoringModelsTemplate.txt found under go/resources folder
//Created @ 2024-10-18 10:32:29.641326 -0500 CDT m=+0.591666917

func BootstrapModels() []ScoringModel {
    var BootstrapModels = []ScoringModel{
        
            { Name: "citi-ts-bv-model",
              MaxScore: 10,
              MinScore: 0,
                Ranges:
                []*Bucket{ 
                    //Bucket
                    { Type: "sloc", Start: "0", End: "int.max",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Deploy to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "10000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "5.00", Outcome: &Outcome{ Score: 0, Recommendation: "Rehost to TKG", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "5.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "10001",
                              End: "10000000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "5.00", Outcome: &Outcome{ Score: 0, Recommendation: "Rehost to TKG", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "5.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "10000001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "5.00", Outcome: &Outcome{ Score: 0, Recommendation: "Rehost to TKG", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "5.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                },
            }, 
            { Name: "con-model",
              MaxScore: 10,
              MinScore: 0,
                Ranges:
                []*Bucket{ 
                    //Bucket
                    { Type: "sloc", Start: "0", End: "1000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "1001", End: "2500",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "2501", End: "5000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "1500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1501",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "5001", End: "10000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Kubernetes Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "10001", End: "int.max",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "10000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "10001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Kubernetes Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_container_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                },
            }, 
            { Name: "Default",
              MaxScore: 10,
              MinScore: 0,
                Ranges:
                []*Bucket{ 
                    //Bucket
                    { Type: "sloc", Start: "0", End: "1000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Deploy to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "1001", End: "2500",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Deploy to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "2501", End: "5000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Deploy to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "1500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1501",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "5001", End: "10000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Deploy to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2", Outcome: &Outcome{ Score: 0, Recommendation: "Rehost to TKG (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "10001", End: "int.max",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to TAS", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "10000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "10001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2.0", Outcome: &Outcome{ Score: 0, Recommendation: "Rehost to TKG (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to TAS (consider modernization)", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                },
            }, 
            { Name: "native-model",
              MaxScore: 10,
              MinScore: 0,
                Ranges:
                []*Bucket{ 
                    //Bucket
                    { Type: "sloc", Start: "0", End: "1000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "1001", End: "2500",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "2501", End: "5000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "1500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1501",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "5001", End: "10000",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Kubernetes Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                    //Bucket
                    { Type: "sloc", Start: "10001", End: "int.max",
                         Ranges:
                        []*Bucket{ 
                            { Type: "raw",
                              Start: "int.min",
                              End: "10", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 10, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "11",
                              End: "100", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.75, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "101",
                              End: "500", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9.5, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "501",
                              End: "1000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "flt.max", Outcome: &Outcome{ Score: 9, Recommendation: "Refactor to Tanzu Application Service", Calculate: false, Expression: "", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "1001",
                              End: "10000", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                            { Type: "raw",
                              Start: "10001",
                              End: "int.max", 
                               Ranges:
                               []*Bucket{ 
                                  { Type: "bv",
                                    Start: "flt.min",
                                    End: "2.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Kubernetes Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "2.01",
                                    End: "8.0", Outcome: &Outcome{ Score: 0, Recommendation: "Refactor to Tanzu Application Service", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                                  { Type: "bv",
                                    Start: "8.01",
                                    End: "flt.max", Outcome: &Outcome{ Score: 0, Recommendation: "Modernization", Calculate: true, Expression: "max_score - log(10,raw_cloud_score)", }, 
                                  }, 
                               },
                            },
                        }, 
                    }, 
                },
            }, 
    }
    return BootstrapModels
}