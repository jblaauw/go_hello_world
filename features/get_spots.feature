Feature: Get spots
    In order to book a spot
    As an employee
    I need to see which spots are available and those that are unavailable

    Scenario: Get all booking information from now until one week from now and there is atleast 1 booking
        Given user is verified as "employee"
        And there is 1 booking
        When a "GET" request is sent to the endpoint "/api/bookings"
        Then the HTTP-response code should be "200"
        And the response should have a list of 1 objects
