// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchlogs_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudWatchLogs_CreateLogGroup() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.CreateLogGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_CreateLogStream() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.CreateLogStream(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteDestination() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteDestinationInput{
		DestinationName: aws.String("DestinationName"), // Required
	}
	resp, err := svc.DeleteDestination(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteLogGroup() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteLogGroup(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteLogStream() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.DeleteLogStream(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteMetricFilterInput{
		FilterName:   aws.String("FilterName"),   // Required
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteMetricFilter(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteRetentionPolicy() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteRetentionPolicyInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteRetentionPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DeleteSubscriptionFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DeleteSubscriptionFilterInput{
		FilterName:   aws.String("FilterName"),   // Required
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteSubscriptionFilter(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DescribeDestinations() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeDestinationsInput{
		DestinationNamePrefix: aws.String("DestinationName"),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.DescribeDestinations(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DescribeLogGroups() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit:              aws.Int64(1),
		LogGroupNamePrefix: aws.String("LogGroupName"),
		NextToken:          aws.String("NextToken"),
	}
	resp, err := svc.DescribeLogGroups(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DescribeLogStreams() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String("LogGroupName"), // Required
		Descending:          aws.Bool(true),
		Limit:               aws.Int64(1),
		LogStreamNamePrefix: aws.String("LogStreamName"),
		NextToken:           aws.String("NextToken"),
		OrderBy:             aws.String("OrderBy"),
	}
	resp, err := svc.DescribeLogStreams(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DescribeMetricFilters() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeMetricFiltersInput{
		LogGroupName:     aws.String("LogGroupName"), // Required
		FilterNamePrefix: aws.String("FilterName"),
		Limit:            aws.Int64(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.DescribeMetricFilters(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_DescribeSubscriptionFilters() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.DescribeSubscriptionFiltersInput{
		LogGroupName:     aws.String("LogGroupName"), // Required
		FilterNamePrefix: aws.String("FilterName"),
		Limit:            aws.Int64(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.DescribeSubscriptionFilters(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_FilterLogEvents() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:  aws.String("LogGroupName"), // Required
		EndTime:       aws.Int64(1),
		FilterPattern: aws.String("FilterPattern"),
		Interleaved:   aws.Bool(true),
		Limit:         aws.Int64(1),
		LogStreamNames: []*string{
			aws.String("LogStreamName"), // Required
			// More values...
		},
		NextToken: aws.String("NextToken"),
		StartTime: aws.Int64(1),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_GetLogEvents() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		EndTime:       aws.Int64(1),
		Limit:         aws.Int64(1),
		NextToken:     aws.String("NextToken"),
		StartFromHead: aws.Bool(true),
		StartTime:     aws.Int64(1),
	}
	resp, err := svc.GetLogEvents(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutDestination() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutDestinationInput{
		DestinationName: aws.String("DestinationName"), // Required
		RoleArn:         aws.String("RoleArn"),         // Required
		TargetArn:       aws.String("TargetArn"),       // Required
	}
	resp, err := svc.PutDestination(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutDestinationPolicy() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutDestinationPolicyInput{
		AccessPolicy:    aws.String("AccessPolicy"),    // Required
		DestinationName: aws.String("DestinationName"), // Required
	}
	resp, err := svc.PutDestinationPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutLogEvents() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{ // Required
			{ // Required
				Message:   aws.String("EventMessage"), // Required
				Timestamp: aws.Int64(1),               // Required
			},
			// More values...
		},
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		SequenceToken: aws.String("SequenceToken"),
	}
	resp, err := svc.PutLogEvents(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutMetricFilterInput{
		FilterName:    aws.String("FilterName"),    // Required
		FilterPattern: aws.String("FilterPattern"), // Required
		LogGroupName:  aws.String("LogGroupName"),  // Required
		MetricTransformations: []*cloudwatchlogs.MetricTransformation{ // Required
			{ // Required
				MetricName:      aws.String("MetricName"),      // Required
				MetricNamespace: aws.String("MetricNamespace"), // Required
				MetricValue:     aws.String("MetricValue"),     // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutMetricFilter(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutRetentionPolicy() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutRetentionPolicyInput{
		LogGroupName:    aws.String("LogGroupName"), // Required
		RetentionInDays: aws.Int64(1),               // Required
	}
	resp, err := svc.PutRetentionPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_PutSubscriptionFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.PutSubscriptionFilterInput{
		DestinationArn: aws.String("DestinationArn"), // Required
		FilterName:     aws.String("FilterName"),     // Required
		FilterPattern:  aws.String("FilterPattern"),  // Required
		LogGroupName:   aws.String("LogGroupName"),   // Required
		RoleArn:        aws.String("RoleArn"),
	}
	resp, err := svc.PutSubscriptionFilter(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleCloudWatchLogs_TestMetricFilter() {
	svc := cloudwatchlogs.New(nil)

	params := &cloudwatchlogs.TestMetricFilterInput{
		FilterPattern: aws.String("FilterPattern"), // Required
		LogEventMessages: []*string{ // Required
			aws.String("EventMessage"), // Required
			// More values...
		},
	}
	resp, err := svc.TestMetricFilter(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}
