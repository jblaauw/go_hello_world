Feature: Book a normal spot as an employee
    In order to work at the office
    As an employee
    I need to be able to reserve a normal spot for myself

    Scenario: Reserve a spot when there is atleast one available
        Given user is verified as "employee" 
        And a "normal" office spot is "available"
        When "employee" books a "normal" spot
        Then save the booking
        And An "success" message is shown. "Your booking has been saved."
        And the HTTP-response code should be "created"

    Scenario: Reserve a spot when the normal spots are unavailable
        Given user is verified as "employee"
        And a "normal" office spot is "unavailable"
        When "employee" books a "normal" spot
        Then An "error" message is shown. "No space is available. Please try other booking spots or contact the admin"
        And the HTTP-response code should be ""

    Scenario: Reserve a spot when the user is not verified as employee
        When "unverified employee" books a "normal" spot
        And a "normal" office spot is "available"
        Then An "error" message is shown. "You are not authorized to book a spot."
        And the HTTP-response code should be "unauthorized"

Feature: Book a normal spot for an employee as an admin
    In order to 
    As 
    I need to be able to 

    Scenario: 
        Given 
        When 
        Then 

Feature: Book an emergency spot for an employee as an admin
    In order to 
    As 
    I need to be able to 

    Scenario: 
        Given 
        When 
        Then 
