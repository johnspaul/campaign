// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": CLI Commands
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --version=v1.3.0

package cli

import (
	"campaign/client"
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateCampaignsCommand is the command line data structure for the create action of campaigns
	CreateCampaignsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteCampaignsCommand is the command line data structure for the delete action of campaigns
	DeleteCampaignsCommand struct {
		// The id of the campaign to be updated
		CampaignID  string
		PrettyPrint bool
	}

	// GetCampaignsCommand is the command line data structure for the get action of campaigns
	GetCampaignsCommand struct {
		// The campaignId
		CampaignID  string
		PrettyPrint bool
	}

	// GetAllCampaignsCommand is the command line data structure for the getAll action of campaigns
	GetAllCampaignsCommand struct {
		// The state
		State       string
		PrettyPrint bool
	}

	// GetAllCampaignExecutionCampaignsCommand is the command line data structure for the getAllCampaignExecution action of campaigns
	GetAllCampaignExecutionCampaignsCommand struct {
		// The campaignId
		CampaignID  string
		PrettyPrint bool
	}

	// GetCampaignExecutionCampaignsCommand is the command line data structure for the getCampaignExecution action of campaigns
	GetCampaignExecutionCampaignsCommand struct {
		// The campaign Id
		CampaignID string
		// The execution Id
		ExecutionID string
		PrettyPrint bool
	}

	// UpdateCampaignsCommand is the command line data structure for the update action of campaigns
	UpdateCampaignsCommand struct {
		Payload     string
		ContentType string
		// The id of the campaign to be updated
		CampaignID  string
		PrettyPrint bool
	}

	// CreateCustomerServiceCommand is the command line data structure for the create action of customerService
	CreateCustomerServiceCommand struct {
		PrettyPrint bool
	}

	// GetLeadCommand is the command line data structure for the get action of lead
	GetLeadCommand struct {
		// the product id
		ProductID   string
		PrettyPrint bool
	}

	// CreateMessagecontentsCommand is the command line data structure for the create action of messagecontents
	CreateMessagecontentsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteMessagecontentsCommand is the command line data structure for the delete action of messagecontents
	DeleteMessagecontentsCommand struct {
		// The message content id
		MessageID   string
		PrettyPrint bool
	}

	// GetMessagecontentsCommand is the command line data structure for the get action of messagecontents
	GetMessagecontentsCommand struct {
		// The message contentid
		MessageID   string
		PrettyPrint bool
	}

	// ListMessagecontentsCommand is the command line data structure for the list action of messagecontents
	ListMessagecontentsCommand struct {
		PrettyPrint bool
	}

	// UpdateMessagecontentsCommand is the command line data structure for the update action of messagecontents
	UpdateMessagecontentsCommand struct {
		Payload     string
		ContentType string
		// The message content id
		MessageID   string
		PrettyPrint bool
	}

	// CreateProductsCommand is the command line data structure for the create action of products
	CreateProductsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// GetProductsCommand is the command line data structure for the get action of products
	GetProductsCommand struct {
		// the product id
		ProductID   string
		PrettyPrint bool
	}

	// CreateSmstrackerserviceCommand is the command line data structure for the create action of smstrackerservice
	CreateSmstrackerserviceCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/"]`,
		Short: `Represents Campaign instance for a Product.`,
		Long: `Represents Campaign instance for a Product.

Payload example:

{
   "activeHours": 1391223088362859816,
   "activeStartHour": 10,
   "activeStartMinute": 53,
   "endDate": 6809827671073780931,
   "executionFrequency": 1,
   "messages": [
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      },
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      },
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      }
   ],
   "pollingInterval": 1932807547281809402,
   "productId": "9y1",
   "startDate": 942289876648515888
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateCustomerServiceCommand)
	sub = &cobra.Command{
		Use:   `customer-service ["/customerservcie/"]`,
		Short: `Mock api for notifying the customer service when a customer replies for an sms`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(CreateMessagecontentsCommand)
	sub = &cobra.Command{
		Use:   `messagecontents ["/messagecontents/"]`,
		Short: `Represents  a message content to be attached to  1 or more campaigns.`,
		Long: `Represents  a message content to be attached to  1 or more campaigns.

Payload example:

{
   "messageContent": "xk8"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(CreateProductsCommand)
	sub = &cobra.Command{
		Use:   `products ["/products/"]`,
		Short: `The details of a product`,
		Long: `The details of a product

Payload example:

{
   "availableLocations": [
      {
         "district": "Corrupti quasi recusandae velit est.",
         "province": "Vel sunt."
      },
      {
         "district": "Corrupti quasi recusandae velit est.",
         "province": "Vel sunt."
      },
      {
         "district": "Corrupti quasi recusandae velit est.",
         "province": "Vel sunt."
      }
   ],
   "clientCode": "Exercitationem inventore qui.",
   "criteria": [
      {
         "comparisonType": "Iste sint voluptatem dolor tempora.",
         "maxValue": 2121481832943976652,
         "minValue": 5207990806549280302,
         "variable": "Odio tempore molestias iusto iure."
      },
      {
         "comparisonType": "Iste sint voluptatem dolor tempora.",
         "maxValue": 2121481832943976652,
         "minValue": 5207990806549280302,
         "variable": "Odio tempore molestias iusto iure."
      },
      {
         "comparisonType": "Iste sint voluptatem dolor tempora.",
         "maxValue": 2121481832943976652,
         "minValue": 5207990806549280302,
         "variable": "Odio tempore molestias iusto iure."
      }
   ],
   "dailyVolume": 6375328854376499111,
   "productCode": "Non nesciunt fugit alias officiis eveniet ratione.",
   "productType": "Eligendi expedita minus voluptatem fugiat."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(CreateSmstrackerserviceCommand)
	sub = &cobra.Command{
		Use:   `smstrackerservice ["/smstrackerservice/"]`,
		Short: `mock api for sms tracker.It creates an sms and registers callback `,
		Long: `mock api for sms tracker.It creates an sms and registers callback 

Payload example:

{
   "callbackApiforResponse": "Veniam nostrum voluptatem quidem.",
   "callbackApiforSend": "Ea deleniti.",
   "campaignId": "Neque vel.",
   "messageContent": "Dolorum natus et fuga.",
   "phoneNumber": 1215192170391213180
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp6 := new(DeleteCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/CAMPAIGNID"]`,
		Short: `Represents Campaign instance for a Product.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(DeleteMessagecontentsCommand)
	sub = &cobra.Command{
		Use:   `messagecontents ["/messagecontents/MESSAGEID"]`,
		Short: `Represents  a message content to be attached to  1 or more campaigns.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `get action`,
	}
	tmp8 := new(GetCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/CAMPAIGNID"]`,
		Short: `Represents Campaign instance for a Product.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(GetLeadCommand)
	sub = &cobra.Command{
		Use:   `lead ["/lead/PRODUCTID"]`,
		Short: `Lead module`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(GetMessagecontentsCommand)
	sub = &cobra.Command{
		Use:   `messagecontents ["/messagecontents/MESSAGEID"]`,
		Short: `Represents  a message content to be attached to  1 or more campaigns.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(GetProductsCommand)
	sub = &cobra.Command{
		Use:   `products ["/products/PRODUCTID"]`,
		Short: `The details of a product`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get-all",
		Short: `Returns all the campaigns`,
	}
	tmp12 := new(GetAllCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/"]`,
		Short: `Represents Campaign instance for a Product.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get-all-campaign-execution",
		Short: `Returns all campaign  execution details.`,
	}
	tmp13 := new(GetAllCampaignExecutionCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/CAMPAIGNID/executions"]`,
		Short: `Represents Campaign instance for a Product.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get-campaign-execution",
		Short: `Returns a specific campaign execution details`,
	}
	tmp14 := new(GetCampaignExecutionCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/CAMPAIGNID/executions/EXECUTIONID"]`,
		Short: `Represents Campaign instance for a Product.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `Returns all the messages.`,
	}
	tmp15 := new(ListMessagecontentsCommand)
	sub = &cobra.Command{
		Use:   `messagecontents ["/messagecontents/"]`,
		Short: `Represents  a message content to be attached to  1 or more campaigns.`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp16 := new(UpdateCampaignsCommand)
	sub = &cobra.Command{
		Use:   `campaigns ["/campaigns/CAMPAIGNID"]`,
		Short: `Represents Campaign instance for a Product.`,
		Long: `Represents Campaign instance for a Product.

Payload example:

{
   "activeHours": 2468065692768091476,
   "activeStartHour": 5545252113718753156,
   "activeStartMinute": 6286374021154434140,
   "endDate": 998879720236526009,
   "executionFrequency": 2406924386454643444,
   "messages": [
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      },
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      },
      {
         "messageId": "7j3jr289vx",
         "percentage": 0.5166763919317907
      }
   ],
   "pollingInterval": 4860372539358858586,
   "startDate": 4982496489640834582,
   "status": 2057154423683357660
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp17 := new(UpdateMessagecontentsCommand)
	sub = &cobra.Command{
		Use:   `messagecontents ["/messagecontents/MESSAGEID"]`,
		Short: `Represents  a message content to be attached to  1 or more campaigns.`,
		Long: `Represents  a message content to be attached to  1 or more campaigns.

Payload example:

{
   "messageContent": "sa"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp17.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the CreateCampaignsCommand command.
func (cmd *CreateCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/campaigns/"
	}
	var payload client.CampaignPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateCampaigns(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteCampaignsCommand command.
func (cmd *DeleteCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/campaigns/%v", url.QueryEscape(cmd.CampaignID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteCampaigns(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var campaignID string
	cc.Flags().StringVar(&cmd.CampaignID, "campaignId", campaignID, `The id of the campaign to be updated`)
}

// Run makes the HTTP request corresponding to the GetCampaignsCommand command.
func (cmd *GetCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/campaigns/%v", url.QueryEscape(cmd.CampaignID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetCampaigns(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var campaignID string
	cc.Flags().StringVar(&cmd.CampaignID, "campaignId", campaignID, `The campaignId`)
}

// Run makes the HTTP request corresponding to the GetAllCampaignsCommand command.
func (cmd *GetAllCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/campaigns/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp18 *float64
	if cmd.State != "" {
		var err error
		tmp18, err = float64Val(cmd.State)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *float64 value", "flag", "--state", "err", err)
			return err
		}
	}
	resp, err := c.GetAllCampaigns(ctx, path, tmp18)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetAllCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var state string
	cc.Flags().StringVar(&cmd.State, "state", state, `The state`)
}

// Run makes the HTTP request corresponding to the GetAllCampaignExecutionCampaignsCommand command.
func (cmd *GetAllCampaignExecutionCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/campaigns/%v/executions", url.QueryEscape(cmd.CampaignID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetAllCampaignExecutionCampaigns(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetAllCampaignExecutionCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var campaignID string
	cc.Flags().StringVar(&cmd.CampaignID, "campaignId", campaignID, `The campaignId`)
}

// Run makes the HTTP request corresponding to the GetCampaignExecutionCampaignsCommand command.
func (cmd *GetCampaignExecutionCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/campaigns/%v/executions/%v", url.QueryEscape(cmd.CampaignID), url.QueryEscape(cmd.ExecutionID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetCampaignExecutionCampaigns(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetCampaignExecutionCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var campaignID string
	cc.Flags().StringVar(&cmd.CampaignID, "campaignId", campaignID, `The campaign Id`)
	var executionID string
	cc.Flags().StringVar(&cmd.ExecutionID, "executionId", executionID, `The execution Id`)
}

// Run makes the HTTP request corresponding to the UpdateCampaignsCommand command.
func (cmd *UpdateCampaignsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/campaigns/%v", url.QueryEscape(cmd.CampaignID))
	}
	var payload client.CampaignUpdatePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateCampaigns(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateCampaignsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var campaignID string
	cc.Flags().StringVar(&cmd.CampaignID, "campaignId", campaignID, `The id of the campaign to be updated`)
}

// Run makes the HTTP request corresponding to the CreateCustomerServiceCommand command.
func (cmd *CreateCustomerServiceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/customerservcie/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateCustomerService(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateCustomerServiceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetLeadCommand command.
func (cmd *GetLeadCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/lead/%v", url.QueryEscape(cmd.ProductID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetLead(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetLeadCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var productID string
	cc.Flags().StringVar(&cmd.ProductID, "productId", productID, `the product id`)
}

// Run makes the HTTP request corresponding to the CreateMessagecontentsCommand command.
func (cmd *CreateMessagecontentsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/messagecontents/"
	}
	var payload client.MessageContentPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateMessagecontents(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateMessagecontentsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteMessagecontentsCommand command.
func (cmd *DeleteMessagecontentsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/messagecontents/%v", url.QueryEscape(cmd.MessageID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteMessagecontents(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteMessagecontentsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var messageID string
	cc.Flags().StringVar(&cmd.MessageID, "messageId", messageID, `The message content id`)
}

// Run makes the HTTP request corresponding to the GetMessagecontentsCommand command.
func (cmd *GetMessagecontentsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/messagecontents/%v", url.QueryEscape(cmd.MessageID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetMessagecontents(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetMessagecontentsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var messageID string
	cc.Flags().StringVar(&cmd.MessageID, "messageId", messageID, `The message contentid`)
}

// Run makes the HTTP request corresponding to the ListMessagecontentsCommand command.
func (cmd *ListMessagecontentsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/messagecontents/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListMessagecontents(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListMessagecontentsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UpdateMessagecontentsCommand command.
func (cmd *UpdateMessagecontentsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/messagecontents/%v", url.QueryEscape(cmd.MessageID))
	}
	var payload client.MessageContentUpdatePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateMessagecontents(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateMessagecontentsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var messageID string
	cc.Flags().StringVar(&cmd.MessageID, "messageId", messageID, `The message content id`)
}

// Run makes the HTTP request corresponding to the CreateProductsCommand command.
func (cmd *CreateProductsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/products/"
	}
	var payload client.ProductPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateProducts(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateProductsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the GetProductsCommand command.
func (cmd *GetProductsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/products/%v", url.QueryEscape(cmd.ProductID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetProducts(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetProductsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var productID string
	cc.Flags().StringVar(&cmd.ProductID, "productId", productID, `the product id`)
}

// Run makes the HTTP request corresponding to the CreateSmstrackerserviceCommand command.
func (cmd *CreateSmstrackerserviceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/smstrackerservice/"
	}
	var payload client.SmsPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateSmstrackerservice(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateSmstrackerserviceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}
