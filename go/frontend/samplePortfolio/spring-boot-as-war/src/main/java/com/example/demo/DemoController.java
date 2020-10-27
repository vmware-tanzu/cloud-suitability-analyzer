package com.example.demo;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController {

    private Data data = new Data();

    @GetMapping("/")
    public String getMessage() {
        return data.getMessage();
    }

    @GetMapping("/data")
    public Data getData() {
        return data;
    }
}
