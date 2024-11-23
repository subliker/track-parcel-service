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
    States() states
}
type states interface{
    MakeParcel() statesmakeParcel
    Register() statesregister
}
type statesmakeParcel interface{
    Name() string
    Recipient() string
    ArrivalAddress() string
    ForecastDate() string
    ForecastDateIncorrectTime() string
    Description() string
    Ready() string
}
type statesregister interface{
    FullName() string
    Email() string
    PhoneNumber() string
    Company() string
    Ready() string
}

type ru_RU_Messages struct{}
func (ru_RU_Messages) OnStartMessage(user_name string) string {
    if user_name == "" {
        return "Привет! Это тестовый бот для менеджеров"
    } else {
        return fmt.Sprintf("Привет, %s! Это тестовый бот для менеджеров", user_name)
    }
}
func (ru_RU_Messages) States() states {
    return ru_RU_states{}
}
type ru_RU_states struct{}
func (ru_RU_states) MakeParcel() statesmakeParcel {
    return ru_RU_statesmakeParcel{}
}
type ru_RU_statesmakeParcel struct{}
func (ru_RU_statesmakeParcel) Name() string {
    return "Давайте добавим вашу посылку! Как вы назовете эту посылку?"
}
func (ru_RU_statesmakeParcel) Recipient() string {
    return "Введите имя получателя"
}
func (ru_RU_statesmakeParcel) ArrivalAddress() string {
    return "Введите адрес получения"
}
func (ru_RU_statesmakeParcel) ForecastDate() string {
    return "Введите предположительную дату доставки"
}
func (ru_RU_statesmakeParcel) ForecastDateIncorrectTime() string {
    return "Введено некорректное время"
}
func (ru_RU_statesmakeParcel) Description() string {
    return "Введите описание для посылки"
}
func (ru_RU_statesmakeParcel) Ready() string {
    return "Посылка была добавлена в систему"
}
func (ru_RU_states) Register() statesregister {
    return ru_RU_statesregister{}
}
type ru_RU_statesregister struct{}
func (ru_RU_statesregister) FullName() string {
    return "Давайте заполним ваш профиль менеджера! Введите отображаемое ФИО менеджера"
}
func (ru_RU_statesregister) Email() string {
    return "Введите отображаемую почту менеджера"
}
func (ru_RU_statesregister) PhoneNumber() string {
    return "Введите отображаемую почту менеджера. Это поле опционально, если не хотите указывать, то нажмите на соответствующую кнопку"
}
func (ru_RU_statesregister) Company() string {
    return "Введите отображаемую компанию менеджера. Это поле опционально, если не хотите указывать, то нажмите на соответствующую кнопку"
}
func (ru_RU_statesregister) Ready() string {
    return "Ваш профиль был успешно создан"
}


