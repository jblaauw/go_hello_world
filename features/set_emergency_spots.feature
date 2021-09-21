Feature: Book an emergency spot for an employee as an admin
    In order to reserve an emergency spot for an employee when the normal spots are fully booked
    As an admin
    I need to be able to book an emergency spot for an employee

    Scenario: An admin books an available emergency spot for an employee
        Given a "emergency" office spot is "available"
        And the request body contains a new "emergency" booking
        And a "POST" request is created for the endpoint "/api/bookings"
        And the request header "Content-Type" is set to "application/json"
        And the user is verified as "admin"
        When the request is sent to the endpoint 
        Then the HTTP-response code should be "created"
        And response body should contain "123123" as it's "BookedBy"
        And response body should contain "EB" as it's "BookingStatus"
