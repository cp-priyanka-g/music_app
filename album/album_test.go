package album

import "testing"

func TestCreate(t *testing.T) {
	expectedData := map[string]string{"album_id": 0, "name": "Sanam", "description": "Recreated old songs ", "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-Justin_Bieber_in_2015.jpg", "is_published": 0, "created_at": "2022-03-21 09:42:41", "updated_at": "2022-03-21 09:42:41", "artist_id": 0}
	actualData, err := Create(expectedData)

	if err != nil {
		t.Errorf("Error occurred while Create Album: %v", err)
	}

	if expectedData != actualData {
		t.Errorf("Expected data %v, but got %v", expectedData, actualData)
	}

}

// func TestInsertAPI(t *testing.T) {
//     // Set up test data
//     expectedData := map[string]string{"id": "1", "name": "Test"}

//     // Call the insert API endpoint
//     actualData, err := insertAPI(expectedData)

//     // Check for errors
//     if err != nil {
//         t.Errorf("Error occurred while calling insert API: %v", err)
//     }

//     // Compare expected and actual data
//     if !reflect.DeepEqual(expectedData, actualData) {
//         t.Errorf("Expected data %v, but got %v", expectedData, actualData)
//     }
// }
