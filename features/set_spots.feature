Feature: Book a normal spot as an employee
    In order to work at the office
    As an employee
    I need to be able to reserve a normal spot for myself

    Scenario: Reserve a spot when there is atleast one available
        Given user is verified as "employee" 
        And a "normal" office spot is "available"
        And the request body contains a new booking
        And the request header content type is set to "application/json"
        When a "POST" request is sent to the endpoint "/api/bookings"
        Then an "success" message is shown. "Your booking has been saved."
        And the HTTP-response code should be "created"
