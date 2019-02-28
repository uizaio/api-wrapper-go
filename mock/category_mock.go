package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

const (
	CategoryBaseURL                       = "/api/public/v3/media/metadata"
	CategoryRelationURL                   = "/api/public/v3/media/entity/related/metadata"
	CategoryId                            = "917ff8b3-1a73-49c6-b989-9604babf785e"
	CategoryId2                           = "1aa8cf57-ab0c-4dca-9a7b-feb2c803905a"
	UpdateCategoryId                      = "3177aac5-8af8-45ec-93cd-1613c832db9e"
	DeleteCategoryId                      = "e96e7331-26f8-48cb-89cb-11b379c33667"
	CreateCategorySuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"917ff8b3-1a73-49c6-b989-9604babf785e\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T04:32:03.261Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"3a04d923-de02-483f-9aa2-4886c0e767d8\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	CreateCategoryFailedResponse          = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"name\",\r\n            \"message\": \"The 'name' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"type\",\r\n            \"message\": \"The 'type' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:43:59.971Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"5e7545b7-929e-482f-bcb1-9b9e9541298c\",\r\n    \"serviceName\": \"api\"\r\n}"
	RetrieveCategorySuccessResponse       = "{\r\n    \"data\": {\r\n        \"id\": \"917ff8b3-1a73-49c6-b989-9604babf785e\",\r\n        \"name\": \"Category name 1\",\r\n        \"description\": \"Category description\",\r\n        \"slug\": \"category-name-1\",\r\n        \"type\": \"folder\",\r\n        \"orderNumber\": 0,\r\n        \"icon\": \"Category icon\",\r\n        \"status\": 1,\r\n        \"createdAt\": \"2019-02-27T04:32:00.000Z\",\r\n        \"updatedAt\": \"2019-02-27T04:32:00.000Z\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T06:16:07.843Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"42dd9ebe-bdab-47ab-8e11-17148a742e38\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	RetrieveUpdateCategoryResponse        = "{\r\n    \"data\": {\r\n        \"id\": \"3177aac5-8af8-45ec-93cd-1613c832db9e\",\r\n        \"name\": \"Update category name\",\r\n        \"description\": \"Update category description\",\r\n        \"slug\": \"update-category-name\",\r\n        \"type\": \"folder\",\r\n        \"orderNumber\": 0,\r\n        \"icon\": \"Update category icon\",\r\n        \"status\": 1,\r\n        \"createdAt\": \"2019-02-27T08:33:44.000Z\",\r\n        \"updatedAt\": \"2019-02-27T08:37:34.000Z\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T08:38:55.323Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"588ebd96-c215-47b5-b187-853e1adc1be2\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	ListCategorySuccessResponse           = "{\r\n    \"data\": [\r\n        {\r\n            \"id\": \"917ff8b3-1a73-49c6-b989-9604babf785e\",\r\n            \"name\": \"Category name 1\",\r\n            \"description\": \"Category description\",\r\n            \"slug\": \"category-name-1\",\r\n            \"type\": \"folder\",\r\n            \"orderNumber\": 0,\r\n            \"icon\": \"Category icon\",\r\n            \"status\": 1,\r\n            \"createdAt\": \"2019-02-27T04:32:00.000Z\",\r\n            \"updatedAt\": \"2019-02-27T04:32:00.000Z\"\r\n        },\r\n        {\r\n            \"id\": \"1aa8cf57-ab0c-4dca-9a7b-feb2c803905a\",\r\n            \"name\": \"sample 3\",\r\n            \"description\": null,\r\n            \"slug\": \"sample-3\",\r\n            \"type\": \"folder\",\r\n            \"orderNumber\": 1,\r\n            \"icon\": null,\r\n            \"status\": 1,\r\n            \"createdAt\": \"2019-02-21T04:00:05.000Z\",\r\n            \"updatedAt\": \"2019-02-21T04:00:05.000Z\"\r\n        }\r\n    ],\r\n    \"metadata\": {\r\n        \"total\": 38,\r\n        \"result\": 2,\r\n        \"page\": 1,\r\n        \"limit\": 2\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T07:23:23.520Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6ac133ef-9194-4671-94f9-1ce99ce205bf\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateCategorySuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"3177aac5-8af8-45ec-93cd-1613c832db9e\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T08:37:36.422Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"ace0b0c6-e7da-47a5-8675-5616b4a66a27\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateCategoryFailedResponse          = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"id\",\r\n            \"message\": \"The 'id' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:49:39.735Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"3ae5d700-baab-496f-8745-6ce37f22a150\",\r\n    \"serviceName\": \"api\"\r\n}"
	DeleteCategorySuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"e96e7331-26f8-48cb-89cb-11b379c33667\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T09:15:17.485Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6f5dd9bb-e483-4d92-87c6-e525786cdc92\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	CreateCategoryRelationSuccessResponse = "{\r\n    \"data\": [\r\n        {\r\n            \"id\": \"b61ab813-cb5b-4299-9b65-ea1ce8111ccd\",\r\n            \"entityId\": \"917ff8b3-1a73-49c6-b989-9604babf785e\",\r\n            \"metadataId\": \"\"\r\n        }\r\n    ],\r\n    \"metadata\": {\r\n        \"total\": 1,\r\n        \"result\": 1,\r\n        \"page\": 1,\r\n        \"limit\": 20\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T09:33:16.962Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"87933461-b7fe-46bd-af3f-eb41fe89f563\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	DeleteCategoryRelationSuccessResponse = "{\r\n    \"data\": [\r\n        {\r\n            \"id\": \"b61ab813-cb5b-4299-9b65-ea1ce8111ccd\",\r\n            \"entityId\": \"917ff8b3-1a73-49c6-b989-9604babf785e\",\r\n            \"metadataId\": \"\"\r\n        }\r\n    ],\r\n    \"metadata\": {\r\n        \"total\": 1,\r\n        \"result\": 1,\r\n        \"page\": 1,\r\n        \"limit\": 20\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T09:37:51.204Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"fb074b87-f7ad-4c6d-ac30-bab88c2a08c4\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
)

var CategoryDataMock = &uiza.CategoryData{
	ID:          *uiza.String(CategoryId),
	Name:        *uiza.String("Category name 1"),
	Type:        uiza.CategoryFolderType,
	Description: *uiza.String("Category description"),
	Slug:        *uiza.String("category-name-1"),
	OrderNumber: *uiza.Int64(0),
	Icon:        *uiza.String("Category icon"),
	Status:      *uiza.Int64(1),
	CreatedAt:   *uiza.String("2019-02-27T04:32:00.000Z"),
	UpdatedAt:   *uiza.String("2019-02-27T04:32:00.000Z"),
}

var CategoryDataMock2 = &uiza.CategoryData{
	ID:          *uiza.String(CategoryId2),
	Name:        *uiza.String("sample 3"),
	Type:        uiza.CategoryFolderType,
	Description: *uiza.String(""),
	Slug:        *uiza.String("sample-3"),
	OrderNumber: *uiza.Int64(1),
	Icon:        *uiza.String(""),
	Status:      *uiza.Int64(1),
	CreatedAt:   *uiza.String("2019-02-21T04:00:05.000Z"),
	UpdatedAt:   *uiza.String("2019-02-21T04:00:05.000Z"),
}

var CategoryUpdateDataMock = &uiza.CategoryData{
	ID:          *uiza.String(UpdateCategoryId),
	Name:        *uiza.String("Update category name"),
	Type:        uiza.CategoryFolderType,
	Description: *uiza.String("Update category description"),
	Slug:        *uiza.String("update-category-name"),
	OrderNumber: *uiza.Int64(0),
	Icon:        *uiza.String("Update category icon"),
	Status:      *uiza.Int64(1),
	CreatedAt:   *uiza.String("2019-02-27T08:33:44.000Z"),
	UpdatedAt:   *uiza.String("2019-02-27T08:37:34.000Z"),
}

var CategoryDeleteDataMock = &uiza.CategoryIDData{
	ID: *uiza.String(DeleteCategoryId),
}

var CategoryDeleteRelationDataMock = []*uiza.CategoryRelationData{
	{
		ID:         *uiza.String("b61ab813-cb5b-4299-9b65-ea1ce8111ccd"),
		EntityId:   *uiza.String("917ff8b3-1a73-49c6-b989-9604babf785e"),
		MetadataId: *uiza.String(""),
	},
}

var CategoryListDataMock = []*uiza.CategoryData{
	CategoryDataMock,
	CategoryDataMock2,
}

type CategoryClientMock struct {
	http.Client
}

func (m *CategoryClientMock) Do(req *http.Request) (*http.Response, error) {
	var typeCategory = uiza.CategoryFolderType
	mockCallTest := []MockData{
		{
			method: "POST",
			path:   CategoryBaseURL,
			params: &uiza.CategoryCreateParams{
				Name:        uiza.String("Category name 1"),
				Type:        &typeCategory,
				Description: uiza.String("Category description"),
				Icon:        uiza.String("Category icon"),
			},
			responseString: CreateCategorySuccessResponse,
		},
		{
			method:         "POST",
			path:           CategoryBaseURL,
			params:         &uiza.CategoryCreateParams{},
			responseString: CreateCategoryFailedResponse,
		},
		{
			method:         "GET",
			path:           CategoryBaseURL,
			params:         &uiza.CategoryIDParams{ID: uiza.String(CategoryId)},
			responseString: RetrieveCategorySuccessResponse,
		},
		{
			method:         "GET",
			path:           CategoryBaseURL,
			params:         &uiza.CategoryIDParams{ID: uiza.String(UpdateCategoryId)},
			responseString: RetrieveUpdateCategoryResponse,
		},
		{
			method: "GET",
			path:   CategoryBaseURL,
			params: &uiza.CategoryListParams{
				Page:  uiza.Int64(1),
				Limit: uiza.Int64(2),
			},
			responseString: ListCategorySuccessResponse,
		},
		{
			method: "PUT",
			path:   CategoryBaseURL,
			params: &uiza.CategoryUpdateParams{
				ID:   uiza.String(UpdateCategoryId),
				Name: uiza.String("Category update name"),
			},
			responseString: UpdateCategorySuccessResponse,
		},
		{
			method:         "PUT",
			path:           CategoryBaseURL,
			params:         &uiza.CategoryUpdateParams{},
			responseString: UpdateCategoryFailedResponse,
		},
		{
			method:         "DELETE",
			path:           CategoryBaseURL,
			params:         &uiza.CategoryIDParams{ID: uiza.String(DeleteCategoryId)},
			responseString: DeleteCategorySuccessResponse,
		},
		{
			method: "POST",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{
				EntityId:    uiza.String(CategoryId),
				MetadataIds: []*string{uiza.String("")},
			},
			responseString: CreateCategoryRelationSuccessResponse,
		},
		{
			method: "DELETE",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{
				EntityId:    uiza.String(CategoryId),
				MetadataIds: []*string{uiza.String("")},
			},
			responseString: DeleteCategoryRelationSuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}
