package vn.northwindondapr.shipping.controller;

import io.cloudevents.CloudEvent;
import lombok.extern.slf4j.Slf4j;
import vn.northwindondapr.shipping.event.OrderCreated;
import vn.northwindondapr.shipping.service.EventService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RequestMapping("/processors")
@RestController
public class ProcessorController {

	EventService eventService;

	@Autowired
	public ProcessorController(EventService eventService) {
		this.eventService = eventService;
	}

	@PostMapping("/order")
	public void subscribe(@RequestBody(required = false) byte[] body) {

		CloudEvent event = eventService.getCloudEvent(body);
		String eventId = event.getId();
		OrderCreated message = eventService.getEventOrderCreated(event);

		log.info("Received eventId: {}, message: {}", eventId, message);
	}

}
