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
    Common() common
    StartMessage() startMessage
    Register() register
    CheckParcel() checkParcel
    Menu() menu
}
type common interface{
    Markup() commonmarkup
}
type commonmarkup interface{
    BtnDontSpecify() string
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
    Ready(name string, email string, phoneNumber string) string
}
type checkParcel interface{
    Main(name string, recipient string, arrivalAddress string, forecastDate string, description string, status string) string
}
type menu interface{
    Main() string
    Markup() menumarkup
}
type menumarkup interface{
    CheckParcel() string
}

type ru_RU_Messages struct{}
func (ru_RU_Messages) Common() common {
    return ru_RU_common{}
}
type ru_RU_common struct{}
func (ru_RU_common) Markup() commonmarkup {
    return ru_RU_commonmarkup{}
}
type ru_RU_commonmarkup struct{}
func (ru_RU_commonmarkup) BtnDontSpecify() string {
    return "Не указывать"
}
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
    return "_Мы делаем отслеживание посылок простым и удобным!_" + "\n" +
        "" + "\n" +
        "_С помощью нашего бота вы сможете:_" + "\n" +
        " • 📦 Отслеживать посылки за секунды" + "\n" +
        " • 🔔 Получать уведомления о каждом изменении статуса" + "\n" +
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
        "_Напишите, пожалуйста, ваше полное имя, чтобы мы могли обращаться к вам по имени. Это и последующие поля будут отображаться публично, так что убедитесь, что вы не против!_" + "\n" +
        "" + "\n" +
        "_Пример: Иванов Иван Иванович_"
}
func (ru_RU_registerpoints) Email() string {
    return "Теперь нам нужно ваш email! 📧" + "\n" +
        "_Это поможет менеджеру ваших посылок поддерживать с вами связь по почте._" + "\n" +
        "" + "\n" +
        "_Пример: ivanov@example.com_"
}
func (ru_RU_registerpoints) PhoneNumber() string {
    return "Теперь давайте добавим ваш номер телефона! 📱" + "\n" +
        "_Этот шаг опционален, но если хотите, чтобы менеджеры могли связаться с вами по телефону — укажите номер._" + "\n" +
        "" + "\n" +
        "_Пример: +7 (999) 123-45-67_"
}
func (ru_RU_registerpoints) Ready(name string, email string, phoneNumber string) string {
    return "Поздравляем, ваш профиль успешно создан! 🎉" + "\n" +
        "Теперь вы можете начать использовать все функции бота. Мы рады приветствовать вас в нашем сообществе! 🙌" + "\n" +
        "Вот что теперь видят менеджеры:" + "\n" +
        fmt.Sprintf("*Имя:* %s", name) + "\n" +
        fmt.Sprintf("*Email:* %s", email) + "\n" +
        fmt.Sprintf("*Телефон* %s", phoneNumber) + "\n" +
        "" + "\n" +
        "_Если нужно изменить какие-то данные, вы всегда можете это сделать в настройках профиля. 👇_"
}
func (ru_RU_Messages) CheckParcel() checkParcel {
    return ru_RU_checkParcel{}
}
type ru_RU_checkParcel struct{}
func (ru_RU_checkParcel) Main(name string, recipient string, arrivalAddress string, forecastDate string, description string, status string) string {
    return "Вот информация по вашей посылке:" + "\n" +
        fmt.Sprintf("Наименование: %s", name) + "\n" +
        fmt.Sprintf("Получатель: %s", recipient) + "\n" +
        fmt.Sprintf("Адрес доставки: %s", arrivalAddress) + "\n" +
        fmt.Sprintf("Ожидаемая дата доставки: %s", forecastDate) + "\n" +
        fmt.Sprintf("Описание: %s", description) + "\n" +
        fmt.Sprintf("Статус: %s", status)
}
func (ru_RU_Messages) Menu() menu {
    return ru_RU_menu{}
}
type ru_RU_menu struct{}
func (ru_RU_menu) Main() string {
    return "Добро пожаловать в наш сервис! 🙌" + "\n" +
        "Здесь вы можете легко отслеживать посылки. Всё, что вам нужно — выбрать нужное действие из меню ниже. 📦" + "\n" +
        "Готовы начать? Выберите одну из *опций*!"
}
func (ru_RU_menu) Markup() menumarkup {
    return ru_RU_menumarkup{}
}
type ru_RU_menumarkup struct{}
func (ru_RU_menumarkup) CheckParcel() string {
    return "📦 Проверить посылку"
}


