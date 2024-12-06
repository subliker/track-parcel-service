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
    Menu() menu
    AddParcel() addParcel
    MyApi() myApi
}
type common interface{
    Markup() commonmarkup
    Errors() commonerrors
}
type commonmarkup interface{
    BtnDontSpecify() string
}
type commonerrors interface{
    Internal() string
    IncorrectInput() string
    Length(min int, max int) string
    TimeFormat() string
    Email() string
    PhoneNumber() string
    AlreadyRegistered() string
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
type menu interface{
    Main() string
    Markup() menumarkup
}
type menumarkup interface{
    AddParcel() string
    MyApi() string
}
type addParcel interface{
    Points() addParcelpoints
}
type addParcelpoints interface{
    Name() string
    Recipient() string
    ArrivalAddress() string
    ForecastDate() string
    Description() string
    Ready(name string, recipient string, arrivalAddress string, forecastDate string, description string, trackNum string) string
}
type myApi interface{
    Main(token string) string
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
func (ru_RU_common) Errors() commonerrors {
    return ru_RU_commonerrors{}
}
type ru_RU_commonerrors struct{}
func (ru_RU_commonerrors) Internal() string {
    return "🚨 Произошла внутренняя ошибка. Пожалуйста, попробуйте позже. Если проблема повторится, свяжитесь с поддержкой."
}
func (ru_RU_commonerrors) IncorrectInput() string {
    return "❌ Неверный ввод!"
}
func (ru_RU_commonerrors) Length(min int, max int) string {
    return fmt.Sprintf("🚫 Некорректная длина: должно быть от %d до %d символов. ✍️ Попробуйте снова!", min, max)
}
func (ru_RU_commonerrors) TimeFormat() string {
    return "⏰ Некорректный формат времени! Пожалуйста, введите дату и время в формате: ДД.ММ.ГГГГ ЧЧ:ММ. Например: 24.12.2024 12:30. Попробуйте снова!"
}
func (ru_RU_commonerrors) Email() string {
    return "❌ Неверный формат почты! Пожалуйста, введите корректный адрес электронной почты. Например: example@mail.com."
}
func (ru_RU_commonerrors) PhoneNumber() string {
    return "❌ Неверный формат номера телефона! Пожалуйста, введите номер в правильном формате. Например: +7 123 456-78-90."
}
func (ru_RU_commonerrors) AlreadyRegistered() string {
    return "⚠️ Ошибка! Этот аккаунт уже зарегистрирован!"
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
        "_Напишите, пожалуйста, ваше полное имя, чтобы мы могли обращаться к вам по имени. Это и последующие поля будут отображаться публично, так что убедитесь, что вы не против!_" + "\n" +
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
        fmt.Sprintf("*Телефон:* %s", phoneNumber) + "\n" +
        fmt.Sprintf("*Компания:* %s", company) + "\n" +
        "" + "\n" +
        "_Если нужно изменить какие-то данные, вы всегда можете это сделать в настройках профиля. 👇_"
}
func (ru_RU_Messages) Menu() menu {
    return ru_RU_menu{}
}
type ru_RU_menu struct{}
func (ru_RU_menu) Main() string {
    return "Добро пожаловать в наш сервис! 🙌" + "\n" +
        "Здесь вы можете легко управлять посылками и добавлять новые. Всё, что вам нужно — выбрать нужное действие из меню ниже. 📦" + "\n" +
        "Готовы начать? Выберите одну из *опций*!"
}
func (ru_RU_menu) Markup() menumarkup {
    return ru_RU_menumarkup{}
}
type ru_RU_menumarkup struct{}
func (ru_RU_menumarkup) AddParcel() string {
    return "📦 Добавить посылку"
}
func (ru_RU_menumarkup) MyApi() string {
    return "🔑 Показать токен"
}
func (ru_RU_Messages) AddParcel() addParcel {
    return ru_RU_addParcel{}
}
type ru_RU_addParcel struct{}
func (ru_RU_addParcel) Points() addParcelpoints {
    return ru_RU_addParcelpoints{}
}
type ru_RU_addParcelpoints struct{}
func (ru_RU_addParcelpoints) Name() string {
    return "Давайте начнем с названия вашей посылки! 📦" + "\n" +
        "_Укажите, как вы хотите назвать вашу посылку. Это поможет вам и получателю ориентироваться среди других посылок._" + "\n" +
        "_Название будет отображаться публично, так что выберите что-то подходящее!_" + "\n" +
        "" + "\n" +
        "_Пример: Ноутбук Xiaomi Redmibook 6600H 16/512gb_"
}
func (ru_RU_addParcelpoints) Recipient() string {
    return "Теперь укажите, кто будет получать посылку! 📧" + "\n" +
        "_Введите имя получателя, чтобы отображалось кому собираетесь доставлять посылку. Это поле будет отображаться публично._" + "\n" +
        "" + "\n" +
        "_Пример: Иванов Иван_"
}
func (ru_RU_addParcelpoints) ArrivalAddress() string {
    return "Теперь укажите адрес, куда нужно доставить посылку! 🏠" + "\n" +
        "_Введите полный адрес доставки. Это поле будет отображаться публично, поэтому убедитесь, что он правильный._" + "\n" +
        "" + "\n" +
        "_Пример: г. Москва, ул. Ленина, 10, кв. 15_"
}
func (ru_RU_addParcelpoints) ForecastDate() string {
    return "Когда вы собираетесь доставить посылку? 📅" + "\n" +
        "_Укажите дату и время, когда посылка предположительно должна быть доставлена. На эти данные будет ориентироваться получатель!._" + "\n" +
        "" + "\n" +
        "_Пример: 24.12.2024 12:30_"
}
func (ru_RU_addParcelpoints) Description() string {
    return "Добавьте описание вашей посылки! ✏️" + "\n" +
        "_Укажите, что именно содержится в посылке, чтобы мы могли помочь вам отслеживать её. Это описание будет публично._" + "\n" +
        "" + "\n" +
        "_Пример: Размеры 350.1 x 242.3 x 14.9 мм Вес 1.8 кг_"
}
func (ru_RU_addParcelpoints) Ready(name string, recipient string, arrivalAddress string, forecastDate string, description string, trackNum string) string {
    return "Ваша посылка успешно добавлена! 🎉" + "\n" +
        "Вот что мы сохранили:" + "\n" +
        fmt.Sprintf("*Название:* %s", name) + "\n" +
        fmt.Sprintf("*Получатель* %s", recipient) + "\n" +
        fmt.Sprintf("*Адрес доставки:* %s", arrivalAddress) + "\n" +
        fmt.Sprintf("*Прогнозируемая дата доставки:* %s", forecastDate) + "\n" +
        fmt.Sprintf("*Описание:* %s", description) + "\n" +
        fmt.Sprintf("Вашей посылке был присужден трек-номер: `%s`", trackNum) + "\n" +
        "" + "\n" +
        "_Если нужно внести изменения, вы всегда можете это сделать в разделе управления посылками. 📝_"
}
func (ru_RU_Messages) MyApi() myApi {
    return ru_RU_myApi{}
}
type ru_RU_myApi struct{}
func (ru_RU_myApi) Main(token string) string {
    return fmt.Sprintf("🔑 Ваш API токен: `%s`", token) + "\n" +
        "" + "\n" +
        "Этот токен можно использовать для обращения к серверу по адресу localhost:8080, чтобы управлять посылками в автоматизированной системе. Пожалуйста, храните токен в безопасности и не передавайте его посторонним лицам."
}


