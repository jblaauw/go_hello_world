Feature: Book a normal spot as an employee
    In order to work at the office
    As an employee
    I need to be able to reserve a normal spot for myself

    Scenario: Reserve a spot when there is atleast one available
        Given a "normal" office spot is "available"
        And the request body contains a new booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        And the user is verified as "employee"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "created"
        And response body should contain "123123" as it's "BookedBy"

    Scenario: Reserve a spot when the normal spots are unavailable
        Given a "normal" office spot is "unavailable"
        And the request body contains a new booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        And the user is verified as "employee"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "service unavailable"
        And response body should contain "There are no spots available." as it's "error"

    Scenario: Reserve a spot when the user is not verified as employee
        Given a "normal" office spot is "available"
        And the request body contains a new booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "unauthorized"
        And response body should contain "You are not authorized to book a spot." as it's "error"
