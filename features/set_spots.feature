Feature: Book a normal spot for an employee as an admin
    In order to let an employee work in the office
    As an admin
    I need to be able to book a normal spot for an employee

    Scenario: Reserve a spot for an employee when there is atleast one available
        Given a "normal" office spot is "available"
        And the request body contains a new "normal" booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        And the user is verified as "admin"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "created"
        And response body should contain "123123" as it's "BookedBy"
        And response body should contain "NB" as it's "BookingStatus"

    Scenario: Reserve a spot for an employee when the normal spots are unavailable
        Given a "normal" office spot is "unavailable"
        And the request body contains a new "normal" booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        And the user is verified as "admin"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "service unavailable"
        And response body should contain "There are no spots available." as it's "error"

    Scenario: Reserve a spot for an employee when the user is not verified as admin
        Given a "normal" office spot is "available"
        And the request body contains a new "normal" booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "unauthorized"
        And response body should contain "You are not authorized to book a spot." as it's "error"
