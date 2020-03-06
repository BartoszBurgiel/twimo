// Import core functionalities from react & react-native library
import React, { useState } from "react";
import {
  View,
  StyleSheet,
  TextInput,
  Text,
  TouchableOpacity
} from "react-native";

// Import color utility
import color from "../../utils/colorPallet";

const SignUpScreen = ({ navigation }) => {
  // State storage for login credentials
  const [userName, setUserName] = useState();
  const [userMail, setUserMail] = useState();
  const [userPwd, setUserPwd] = useState();
  let userData = [];

  // State storage for socket
  const [socketMessages, setSocketMessages] = useState();

  // Submit user credentials
  const initSignUp = () => {
    userData.push(userName, userMail, userPwd);
    connectToSocket(userData);
    userData = [];
  };

  // Connect to the websocket
  const connectToSocket = arr => {
    // Socket url
    const socket = new WebSocket("ws://localhost:8080/signup");

    //Open socket and send data
    socket.onopen = () => {
      // Loop array
      arr.map(item => {
        socket.send(item);
      });
      //Receive validation hash
      while (!socketMessages) {
        socket.onmessage = response => {
          let data = JSON.parse(response.data);
          setSocketMessages(data);
        };
      }
      // Log validation
      console.log(socketMessages);
    };
  };

  // Render component
  return (
    <View style={styles.container}>
      <Text style={styles.header}>Sign up</Text>
      <Text style={styles.label}>Name</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="e.g. Luke Skywalker"
        placeholderTextColor={color.brand.dark}
        onChangeText={text => setUserName(text)}
      />
      <Text style={styles.label}>Email</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        placeholder="e.g. luke@jedi-mail.com"
        placeholderTextColor={color.brand.dark}
        onChangeText={text => setUserMail(text)}
      />
      <Text style={styles.label}>Password</Text>
      <TextInput
        style={styles.input}
        autoCompleteType="off"
        secureTextEntry={true}
        placeholder="e.g. RedLightSaber#69"
        placeholderTextColor={color.brand.dark}
        onChangeText={text => setUserPwd(text)}
      />
      <TouchableOpacity
        style={styles.buttonPrimary}
        onPress={() => {
          navigation.navigate("List");
          initSignUp();
        }}
      >
        <Text style={styles.buttonPrimaryText}>Let's go!</Text>
      </TouchableOpacity>
    </View>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  },
  header: {
    fontSize: 30,
    fontWeight: "bold",
    marginBottom: 20,
    color: color.brand.normal
  },
  label: {
    width: "70%",
    textAlign: "left",
    textTransform: "uppercase",
    fontWeight: "bold",
    color: color.black
  },
  input: {
    height: 50,
    width: "70%",
    marginVertical: 10,
    paddingHorizontal: 10,
    borderRadius: 12,
    backgroundColor: color.brand.light,
    color: color.brand.dark
  },
  buttonPrimary: {
    backgroundColor: color.brand.normal,
    height: 50,
    width: "50%",
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 20
  },
  buttonPrimaryText: {
    color: color.white,
    textTransform: "uppercase",
    fontWeight: "bold",
    fontSize: 20
  }
});

// Export Screen
export default SignUpScreen;
