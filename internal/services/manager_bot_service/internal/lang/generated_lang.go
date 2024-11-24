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
    StartMessage() startMessage
    Register() register
    States() states
}
type startMessage interface{
    Head(user_name string) string
    Main() string
    Markup() startMessagemarkup
}
type startMessagemarkup interface{
    Register() string
}
type register interface{
    Points() registerpoints
}
type registerpoints interface{
    FullName() string
    Email() string
    PhoneNumber() string
    Company() string
    Ready(name string, email string, phoneNumber string, company string) string
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
    Ready(trackNumber string) string
}
type statesregister interface{
    FullName() string
    Email() string
    PhoneNumber() string
    Company() string
    Ready() string
}

type ru_RU_Messages struct{}
func (ru_RU_Messages) StartMessage() startMessage {
    return ru_RU_startMessage{}
}
type ru_RU_startMessage struct{}
func (ru_RU_startMessage) Head(user_name string) string {
    if user_name == "" {
        return "✨ Добро пожаловать! ✨"
    } else {
        return fmt.Sprintf("✨ Добро пожаловать, %s! ✨", user_name)
    }
}
func (ru_RU_startMessage) Main() string {
    return "_Мы упрощаем работу менеджеров с посылками! С помощью нашего бота вы сможете:_" + "\n" +
        "" + "\n" +
        " • 📦 Легко добавлять и отслеживать посылки" + "\n" +
        " • 🗺️ Добавлять чекпоинты с подробностями" + "\n" +
        " • 🔍 Быстро находить нужную информацию" + "\n" +
        "" + "\n" +
        "_Начните с регистрации, чтобы открыть все возможности бота!_" + "\n" +
        "*Кнопка*: \"✅ Зарегистрироваться\""
}
func (ru_RU_startMessage) Markup() startMessagemarkup {
    return ru_RU_startMessagemarkup{}
}
type ru_RU_startMessagemarkup struct{}
func (ru_RU_startMessagemarkup) Register() string {
    return "✅ Зарегистрироваться"
}
func (ru_RU_Messages) Register() register {
    return ru_RU_register{}
}
type ru_RU_register struct{}
func (ru_RU_register) Points() registerpoints {
    return ru_RU_registerpoints{}
}
type ru_RU_registerpoints struct{}
func (ru_RU_registerpoints) FullName() string {
    return "Давайте начнем с вашего имени! 📝" + "\n" +
        "_Напишите, пожалуйста, ваше полное имя, чтобы мы могли обращаться к вам по имени. Это и последующие поля будет отображаться публично, так что убедитесь, что вы не против!_" + "\n" +
        "" + "\n" +
        "_Пример: Иванов Иван Иванович_"
}
func (ru_RU_registerpoints) Email() string {
    return "Теперь нам нужно ваш email! 📧" + "\n" +
        "_Это поможет вашим получателям поддерживать с вами связь по почте._" + "\n" +
        "" + "\n" +
        "_Пример: ivanov@example.com_"
}
func (ru_RU_registerpoints) PhoneNumber() string {
    return "Теперь давайте добавим ваш номер телефона! 📱" + "\n" +
        "_Этот шаг опционален, но если хотите, чтобы получатели могли связаться с вами по телефону — укажите номер._" + "\n" +
        "" + "\n" +
        "_Пример: +7 (999) 123-45-67_"
}
func (ru_RU_registerpoints) Company() string {
    return "Укажите, пожалуйста, название вашей компании! 🏢" + "\n" +
        "_Это поле опционально, но если у вас есть компания, то добавьте ее сюда. Название компании будет видно публично._" + "\n" +
        "" + "\n" +
        "_Пример: ООО 'ТехноПарт'_"
}
func (ru_RU_registerpoints) Ready(name string, email string, phoneNumber string, company string) string {
    return "Поздравляем, ваш профиль успешно создан! 🎉" + "\n" +
        "Теперь вы можете начать использовать все функции бота. Мы рады приветствовать вас в нашем сообществе! 🙌" + "\n" +
        "Вот что теперь видят получатели:" + "\n" +
        fmt.Sprintf("*Имя:* %s", name) + "\n" +
        fmt.Sprintf("*Email:* %s", email) + "\n" +
        fmt.Sprintf("*Телефон* %s", phoneNumber) + "\n" +
        fmt.Sprintf("*Компания:* %s", company) + "\n" +
        "" + "\n" +
        "_Если нужно изменить какие-то данные, вы всегда можете это сделать в настройках профиля. 👇_"
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
func (ru_RU_statesmakeParcel) Ready(trackNumber string) string {
    return fmt.Sprintf("Посылка была добавлена в систему. Ее трек номер: %s", trackNumber)
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


