package controllers

import (
    "context"

    "cloud.google.com/go/firestore"
    "firebase.google.com/go"
    "google.golang.org/api/option"
)

type Skill struct {
    Category  Category  `json:"category"`
    CreatedAt time.Time `json:"created_at"`
    Detail    string    `json:"detail"`
    Duration  int64     `json:"duration"`
    Name      string    `json:"name"`
    SelfEval  int64     `json:"self_evaluation"`
    Term      string    `json:"term"`
}

type Category struct {
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
}

type Skills *[]Skill

type SkillController struct {
    Interactor usecase.SkillInteractor
}

func (controller *SkillController) Index(c Context) {

    ctx := context.Background()

    client, err := firebaseInit(ctx)
    if err != nil {

    }

    data := client.Collection(term).Documents(ctx)

    docs, err := data.GetAll()
    if err != nil {

    }

    // 配列の初期化
    skills := make([]*Skill, 0)
    for _, doc := range docs {
        // 構造体の初期化
        s := new(Skill)
        // 構造体にFirestoreのデータをセット
        mapToStruct(doc.Data(), &s)
        // 配列に構造体をセット
        skills = append(skills, s)
    }

    defer client.Close()

    c.JSON(200, skills)

}

// func (controller *SkillController) Index(c Context) {
//     skills, error := controller.Interactor.Skills()
//     if err != nil {
//         c.JSON(500, NewError(err))
//         return
//     }
//     c.JSON(200, skills)
// }


func firebaseInit(ctx *gin.Context) (*firestore.Client, error) {

    sa := option.WithCredentialsFile("../../../portfolio-firebase-adminsdk.json")

    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {

    }

    client, err := app.Firestore(ctx)
    if err != nil {

    }

    return client, nil

}

// mapから構造体に変換を行う
func mapToStruct(m map[string]interface{}, val interface{}) error {
    tmp, err := json.Marshal(m)
    if err != nil {
        return err
    }
    err = json.Unmarshal(tmp, val)
    if err != nil {
        return err
    }
    return nil
}

