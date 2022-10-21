package structs

type NotePayload struct {
	Header  string `dynamodbav:"header"`
	Content string `dynamodbav:"content"`
}

type Note struct {
	ID string `dynamodbav:"id"`
	NotePayload
}
