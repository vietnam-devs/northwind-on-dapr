package vn.northwindondapr.shipping.controller;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import vn.northwindondapr.shipping.pojo.Subscription;

@RestController
public class HomeController {

	@GetMapping("/")
	public String index() {
		return "Greetings from Spring Boot!";
	}

	@GetMapping("/dapr/subscribe")
	public List<Subscription> subscribe() {
		Subscription subscription = new Subscription("pubsub", "order", "/processors/order");
		return List.of(subscription);
	}

}
