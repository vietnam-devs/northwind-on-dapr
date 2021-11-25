package vn.northwindondapr.shipping.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.cloudevents.CloudEvent;
import io.cloudevents.core.data.PojoCloudEventData;
import io.cloudevents.core.format.EventFormat;
import io.cloudevents.core.provider.EventFormatProvider;
import io.cloudevents.jackson.JsonFormat;
import io.cloudevents.jackson.PojoCloudEventDataMapper;
import vn.northwindondapr.shipping.event.OrderCreated;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import static io.cloudevents.core.CloudEventUtils.mapData;


@Service
public class EventService {
    private ObjectMapper objectMapper;

    @Autowired
    public EventService(ObjectMapper objectMapper) {
        this.objectMapper = objectMapper;
    }

    public OrderCreated getEventOrderCreated(CloudEvent event) {
        PojoCloudEventData<OrderCreated> cloudEventData = mapData(
                event,
                PojoCloudEventDataMapper.from(objectMapper, OrderCreated.class)
        );
        return cloudEventData.getValue();
    }

    public CloudEvent getCloudEvent(byte[] body) {
        EventFormat format = EventFormatProvider
                .getInstance()
                .resolveFormat(JsonFormat.CONTENT_TYPE);
        return format.deserialize(body);
    }
}
