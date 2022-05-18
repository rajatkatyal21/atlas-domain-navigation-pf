##### What instrumentation this service would need to ensure its observability and operational transparency?
````aidl
Integrating the tools like datadog with the app would provide great level of monitoring for the system. 
We should have traceId added to the header requests to make sure that each request should be traceable and we 
can use these traceIds for debugging using the app logs.
We should also monitor the container level monitoring with the tools like Graphana.
````
##### Why throttling is useful (if it is)? How would you implement it here?
```aidl
Throttling is useful in cases when we think that the system is wont be able to handle the traffic after an certain extent.
It can help in making sure that the System is not broken when high traffic comes to system and also prevents the components like 
DB.

We can apply the throtteling using the BUS layer/Load balancer before the app.
Another unorthodox way could be use some message queue between the client and app server. This can buffer the message in queue and make sure that the client wont have any impact with the high 
traffic. 

``` 
##### What we have to change to make DNS be able to service several sectors at the same time?
```aidl
We just need to add the sector ID in the environment variable and change the endpoint of the API to have sector identifier.
The load balancer needs to add the routing based on the sectorId id to the respective container.  

``` 
##### Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by allowing cargo ships of MomCorp to use DNS. The only issue is - MomCorp software expects loc value in location field, but math stays the same. How would you approach this? What’s would be your implementation strategy?
```aidl
I will create a new field in the API response called location. Now the API will respond with loc and location with the same value.
Example : {
    "loc": "1234",
    "location": "1234"
}
```

##### Atlas Corp mathematicians made another breakthrough and now our navigation math is even better and more accurate, so we started producing a new drone model, based on new math. How would you enable scenario where DNS can serve both types of clients? 
````aidl
We can create a new version and deploy that a new a conatiner with the updated location calculation. So, now the old client will route to the old app like v1 and new clients will go to the v2 version.  
````
##### In general, how would you separate technical decision to deploy something from business decision to release something?

````aidl
IMO, we should have different environment in production. Like Pre-release (Canary) deployments where the features can be tested and when the business approves it or gives a go ahead then the should 
release the fetaure/product to the prouction and change the routing to use the new image-> containers
The containers of canary and prouction environment should use the same componemts like storage service, DB, Messaging Q brokers etc. 

````
