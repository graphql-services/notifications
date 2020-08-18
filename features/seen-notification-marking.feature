Feature: User should be able to mark one or more notifications as seen=true

    Background: We have test notifications
        Given I send query:
            """
            mutation {
            deleteAllNotifications
            n1: createNotification(input:{message:"test1",principal:"aaa",reference:"User",referenceID:"john.doe",date:"2019-01-01T15:00:00Z",seen:false}) { id }
            n2: createNotification(input:{message:"test2",principal:"aaa",reference:"User",referenceID:"example",date:"2019-01-01T15:00:00Z",seen:false}) { id }
            n3: createNotification(input:{id:"test3",message:"test3",principal:"aaa",reference:"User",referenceID:"john.doe",date:"2019-01-01T15:00:00Z",seen:false}) { id }
            }
            """
        And the error should be:
            """
            null
            """

    Scenario: Fetching notifications should return valid result
        When I send query:
            """
            query {
            notifications(sort:[{message:ASC}]) { items { message reference referenceID seen } }
            }
            """
        Then the response should be:
            """
            {
                "notifications": {
                    "items": [
                        {
                            "message": "test1",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": false
                        },
                        {
                            "message": "test2",
                            "reference": "User",
                            "referenceID": "example",
                            "seen": false
                        },
                        {
                            "message": "test3",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": false
                        }
                    ]
                }
            }
            """

    Scenario: Marking notification should work properly
        When I send query:
            """
            mutation {
            seenNotification(id:"test3") { id }
            }
            """
        Then I send query:
            """
            query {
            notifications(sort:[{message:ASC}]) { items { message reference referenceID seen } }
            }
            """
        Then the response should be:
            """
            {
                "notifications": {
                    "items": [
                        {
                            "message": "test1",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": false
                        },
                        {
                            "message": "test2",
                            "reference": "User",
                            "referenceID": "example",
                            "seen": false
                        },
                        {
                            "message": "test3",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": true
                        }
                    ]
                }
            }
            """


    Scenario: Batch marking notifications should work properly
        When I send query:
            """
            mutation {
            seenNotifications(reference:"User",principal:"aaa",referenceID:"john.doe")
            }
            """
        Then I send query:
            """
            query {
            notifications(sort:[{message:ASC}]) { items { message reference referenceID seen } }
            }
            """
        Then the response should be:
            """
            {
                "notifications": {
                    "items": [
                        {
                            "message": "test1",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": true
                        },
                        {
                            "message": "test2",
                            "reference": "User",
                            "referenceID": "example",
                            "seen": false
                        },
                        {
                            "message": "test3",
                            "reference": "User",
                            "referenceID": "john.doe",
                            "seen": true
                        }
                    ]
                }
            }
            """