# test/feature/login.feature
Feature: Login

  Scenario: Successful login
    Given I have the correct email and password
    When I submit the login form
    Then I should receive a successful response
