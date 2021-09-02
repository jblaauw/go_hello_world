Feature: Set normal spot
    In order to work at the office
    As an employee
    I need to be able to reserve a normal spot

    Scenario: Reserve a spot when there is atleast one available
        Given User input is validated as an employee or admin and office spots are available 
        When Employee tries to book the spot
        Then Add the spot object in the database and success message is shown

    Scenario: Reserve a spot when the normal spots are unavailable
        Given User input is validated as employee and office reservation spots are unavailable
        When Employee tries to book the spot
        Then An error message is shown. "No space is available. Please try other booking spots or contact the admin".

    Scenario: Reserve a spot when the user is not verified as employee
        Given User input is not validated as employee
        When User tries to book the spot
        Then An error message is shown. "You do not have the permission to book a spot.".
