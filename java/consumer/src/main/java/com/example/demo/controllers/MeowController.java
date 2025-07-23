package com.example.demo.controllers;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.example.demo.services.ProducerService;

@RestController
public class MeowController {
    private final ProducerService producerService;

    public MeowController(ProducerService producerService) {
        this.producerService = producerService;
    }

    @PostMapping("/meow")
    public String sendMsg(@RequestParam(name = "name", defaultValue = "undefined") String name) {
        producerService.sendMessage("test-topic-1", name);
        return "Send succesfully" + name;
    }
}
