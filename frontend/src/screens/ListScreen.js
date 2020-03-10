// Import core functionalities from react & react-native library
import React, { useState, useEffect, useRef } from "react";
import { View, FlatList, StyleSheet } from "react-native";

// Import custom components
import LocationCard from "../components/LocationCard";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define the ListScreen
const ListScreen = ({ navigation }) => {
  // Web socket state
  const [socketMessages, setSocketMessages] = useState();

  // Web socket connection
  useEffect(() => {
    // Define socket
    let socket = new WebSocket("ws://localhost:8080/listscreen");

    // Fetch JSON sorted by rating
    socket.onopen = () => {
      socket.send("rating");
      socket.onmessage = response => {
        let data = JSON.parse(response.data);
        setSocketMessages(data);
      };
    };

    // Close socket connection if data is fetched
    if (socketMessages) {
      socket.close();
    }
  }, []);

  // Iterate over the given nodes and display them as LocationCards in a FlatList
  return (
    <View style={styles.container}>
      <FlatList
        data={socketMessages}
        keyExtractor={item => item.ID}
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
