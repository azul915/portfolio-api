package main

import (

    "context"
    "encoding/json"
    "log"
    "time"

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

func getServerSideSkill() ([]*Skill, error) {

    ctx := context.Background()
    client, err := firebaseInit(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ取得
    data := client.Collection("serverside").Documents(ctx)

    // 全ドキュメント取得
    docs, err := data.GetAll()
    if err != nil {
        log.Fatalln(err)
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

    return skills, err
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

func setServerSideSkill() error {

    ctx := context.Background()
    client, err := firebaseInit(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ追加
    collection := client.Collection("serverside")
    doc := collection.Doc("Go")
    _, err = doc.Set(ctx, Skill {
        Category: Category {
            ID: 0,
            Name: "言語",
        },
        CreatedAt: time.Now(),
        Detail: "SpringBootでAPIコンテナを実装したことがある。",
        Duration: 3,
        Name: "Kotlin",
        SelfEval: 4,
        Term: "serverside",
    })

    if err != nil {
        log.Fatalln(err)
    }

    defer client.Close()

    return err

}