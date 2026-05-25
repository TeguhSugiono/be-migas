#include <WiFi.h>
#include <HTTPClient.h>
#include <ArduinoOTA.h>

//const char* ssid = "ahla_B10";
//const char* password = "Ahlanuha1903anak2";

const char* ssid = "smaihbsx1-3";
const char* password = "12345678";

// URL API GOLANG
const char* serverName = "http://192.168.1.100:2211/migas/api/v1/gas";

const char* deviceId = "GAS-29FXHH";
const char* deviceToken = "1P2CQSGNGA0XZSD15NEQ";

// PIN
#define MQ6_PIN 34
#define BUZZER 26
#define RELAY 27

int batasGas = 1500;

void setup()
{
  Serial.begin(115200);

  pinMode(BUZZER, OUTPUT);
  pinMode(RELAY, OUTPUT);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi Connected");
  Serial.println(WiFi.localIP());

  // OTA
  ArduinoOTA.setHostname(deviceId);
  ArduinoOTA.setPassword("1234567890");

  ArduinoOTA.begin();

  Serial.println("OTA READY");
}

unsigned long lastSend = 0;
unsigned long interval = 3000; // 3 detik

// 20 request / menit
// 1200 request / jam
// 28800 request / hari

void loop()
{
  ArduinoOTA.handle();

  if(millis() - lastSend > interval)
  {
    lastSend = millis();

    //int gasValue = 200;

    // nanti kalau sensor sudah dipasang:
    int gasValue = analogRead(MQ6_PIN);

    Serial.print("Gas Value: ");
    Serial.println(gasValue);

    bool statusGas = false;
    String alarmStatus = "AMAN";

    if(gasValue > batasGas)
    {
      statusGas = true;
      alarmStatus = "BAHAYA";

      digitalWrite(BUZZER, HIGH);
      digitalWrite(RELAY, HIGH);
    }
    else
    {
      digitalWrite(BUZZER, LOW);
      digitalWrite(RELAY, LOW);
    }

    kirimData(statusGas, alarmStatus);
  }
}

void kirimData(bool statusGas, String alarmStatus)
{
  if(WiFi.status() == WL_CONNECTED)
  {
    HTTPClient http;

    Serial.println("CONNECT API...");
    Serial.println(serverName);

    http.begin(serverName);

    // TIMEOUT 3 DETIK
    http.setTimeout(3000);

    http.addHeader("Content-Type", "application/json");

    String jsonData = "{";
    jsonData += "\"device_id\":\"";
    jsonData += deviceId;
    jsonData += "\",";

    jsonData += "\"device_token\":\"";
    jsonData += deviceToken;
    jsonData += "\",";

    jsonData += "\"status_gas\":";
    jsonData += (statusGas ? "true" : "false");
    jsonData += ",";

    jsonData += "\"alarm_status\":\"";
    jsonData += alarmStatus;
    jsonData += "\",";

    jsonData += "\"wifi_status\":true";
    jsonData += "}";

    Serial.println("JSON:");
    Serial.println(jsonData);

    int responseCode = http.POST(jsonData);

    Serial.print("Response Code: ");
    Serial.println(responseCode);

    // CEK HASIL REQUEST
    if(responseCode > 0)
    {
      Serial.println("KIRIM SUKSES");

      String response = http.getString();

      Serial.println(response);
    }
    else
    {
      Serial.println("BACKEND OFFLINE");

      Serial.print("HTTP ERROR: ");
      Serial.println(http.errorToString(responseCode));
    }

    http.end();
  }
  else
  {
    Serial.println("WIFI DISCONNECTED");
  }
}
