Feature: Get spots
    In order to book a spot
    As an employee
    I need to see which spots are available and those that are unavailable

    Scenario: Get all booking information from now until one week from now and there is atleast 1 booking
        Given a "GET" request is created for the endpoint "/api/bookings"
        And the user is verified as "employee"
        And there is 1 booking
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "success"
        And the response should have a list of 1 objects
