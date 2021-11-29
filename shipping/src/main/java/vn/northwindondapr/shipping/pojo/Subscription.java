package vn.northwindondapr.shipping.pojo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Subscription {
	private String pubsubname;
	private String topic;
	private String route;
}