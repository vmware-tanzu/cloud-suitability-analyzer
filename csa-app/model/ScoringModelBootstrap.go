/*******************************************************************************
 * Copyright (c) 2018 - Present VMWare, Inc. All Rights Reserved.
 ******************************************************************************/

package model

//Created By BootstrapScoringModelsTemplate.txt found under go/resources folder
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD:go/model/ScoringModelBootstrap.go
<<<<<<< HEAD
<<<<<<< HEAD
//Created @ 2022-10-03 17:56:47.78384 -0400 EDT m=+0.283689696
=======
//Created @ 2022-06-27 13:51:59.901432 -0500 CDT m=+0.190101183
>>>>>>> 21df02c (Title: Add newly revised rules)
=======
//Created @ 2022-06-27 14:22:18.591091 -0500 CDT m=+0.234998174
>>>>>>> 35d7036 (Title: Add new rules)
=======
//Created @ 2022-06-27 15:42:06.015881 -0500 CDT m=+0.211386837
>>>>>>> dacc802 (Title: Create a local build):csa-app/model/ScoringModelBootstrap.go
=======
//Created @ 2022-06-27 19:36:34.536928 -0500 CDT m=+0.199993582
>>>>>>> 32e5db1 (Title:  Fix local build with not github path)
=======
//Created @ 2022-06-27 20:06:27.393071 -0500 CDT m=+0.181195651
>>>>>>> 0405f37 (Title:  Fix some binning errors)
=======
//Created @ 2022-06-28 08:54:27.413137 -0500 CDT m=+0.181809344
>>>>>>> ae9093c (Title: Fix errors in rules/bins)
=======
//Created @ 2022-06-28 10:47:48.975553 -0500 CDT m=+0.173371515
>>>>>>> 92324ac (Title: standard tags/categories)
=======
//Created @ 2022-06-30 11:37:53.399988 -0500 CDT m=+0.215968101
>>>>>>> 64ece2d (Title: Remove contains rules where possible)

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
    }
    return BootstrapModels
}