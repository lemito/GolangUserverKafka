package com.example.demo.services;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class ConsumerService {

    @KafkaListener(topics = "test-topic-1", groupId = "test-group-1")
    public void listen(String message) {
        System.out.println("Received Message: " + message);
    }

}
