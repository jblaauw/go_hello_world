Feature: Get spot
    In order to book a spot
    As an employee
    I need to see which spots are available and those that are unavailable

    Scenario: Get 1 specific booking information
        Given there is 1 booking
        And a "GET" request is created for the endpoint "/api/me/bookings/1"
        And the user is verified as "employee"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "success"
        And response body should contain "NB" as it's "BookingStatus"

    Scenario: Get 1 specific booking information that does not exist
        Given there is 1 booking
        And a "GET" request is created for the endpoint "/api/me/bookings/17"
        And the user is verified as "employee"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "not found"
        And response body should contain "The booking was not found." as it's "error"
