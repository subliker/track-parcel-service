/** Code generated using https://github.com/MrNemo64/go-n-i18n 
 * Any changes to this file will be lost on the next tool run */

package lang

import (
    "fmt"
    "strings"
)

func MessagesFor(tag string) (Messages, bool) {
    switch strings.ReplaceAll(tag, "_", "-") {
    case "ru-RU":
        return ru_RU_Messages{}, true
    }
    return nil, false
}

func MessagesForMust(tag string) Messages {
    switch strings.ReplaceAll(tag, "_", "-") {
    case "ru-RU":
        return ru_RU_Messages{}
    }
    panic(fmt.Errorf("unknwon language tag: " + tag))
}

func MessagesForOrDefault(tag string) Messages {
    switch strings.ReplaceAll(tag, "_", "-") {
    case "ru-RU":
        return ru_RU_Messages{}
    }
    return ru_RU_Messages{}
}

type Messages interface{
    OnStartMessage(user_name string) string
    OnAddParcel() string
}

type ru_RU_Messages struct{}
func (ru_RU_Messages) OnStartMessage(user_name string) string {
    if user_name == "" {
        return "Привет! Это тестовый бот для менеджеров"
    } else {
        return fmt.Sprintf("Привет, %s! Это тестовый бот для менеджеров", user_name)
    }
}
func (ru_RU_Messages) OnAddParcel() string {
    return "Давайте добавим вашу посылку! Как вы назовете эту посылку?"
}


