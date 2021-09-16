Feature: Book a normal spot as an employee
    In order to work at the office
    As an employee
    I need to be able to reserve a normal spot for myself

    Scenario: Reserve a spot when there is atleast one available
        Given a "normal" office spot is "available"
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header content type is set to "application/json"
        And the user is verified as "employee" 
        And the request body contains a new booking
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "created"
        And response body should contain "123123" as it's "BookedBy"
