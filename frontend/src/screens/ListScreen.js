// Import core functionalities from react & react-native library
import React, { useState, useEffect, useRef } from "react";
import { View, FlatList, StyleSheet } from "react-native";

// Import custom components
import LocationCard from "../components/LocationCard";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define the ListScreen
const ListScreen = ({ navigation }) => {
  // Enable an mutable state
  const [locationData, setLocationData] = useState(0);

  // Fetch Location data from web API and store it in state
  useEffect(() => {
    fetch("http://127.0.0.1:5500/frontend/data/locations.json")
      .then(response => response.json())
      .then(data => setLocationData(data));
  });

  // Web socket
  const [socketMessages, setSocketMessages] = useState();

  useEffect(() => {
    let socket = new WebSocket("ws://localhost:8080/listscreen");
    socket.onopen = () => {
      socket.send("price");
      socket.onmessage = response => {
        setSocketMessages(response.data);
      };
    };
    console.log(socketMessages);

    if (socketMessages) {
      socket.close();
    }
  }, [socketMessages]);

  // Iterate over the given nodes and display them as LocationCards in a FlatList
  return (
    <View style={styles.container}>
      <FlatList
        data={socketMessages}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => (
          <LocationCard data={item} navigation={navigation} />
        )}
      />
    </View>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    width: "100%",
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  }
});

// Export the above defined screen
export default ListScreen;
