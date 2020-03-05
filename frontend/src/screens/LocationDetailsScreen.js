// Import core functionality dependencies from react & react-native
import React, { useState, useEffect } from "react";
import {
  View,
  SafeAreaView,
  FlatList,
  Text,
  StyleSheet,
  Image
} from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Import custom built components
import RatingStars from "../components/RatingStars";
import FeatureIcons from "../components/FeatureIcons";
import ReviewCard from "../components/ReviewCard";

// Define LocationDetailsScreen
const LocationDetailsScreen = props => {
  // Grab locationData from the props and store it in data const
  const data = props.navigation.state.params.locationData;
  // Extract locationID from data prop
  let locationID = data.ID;

  // Web socket state
  const [socketReviews, setSocketReviews] = useState();

  // Web socket connection
  useEffect(() => {
    // Define socket
    let socket = new WebSocket(`ws://localhost:8080/location/${locationID}`);

    // Fetch JSON sorted by rating
    socket.onmessage = response => {
      let data = JSON.parse(response.data);
      setSocketReviews(data.Comments);
    };

    // Close socket connection if data is fetched
    if (socketReviews) {
      socket.close();
    }
  }, [socketReviews]);

  // Render LocationDetailScreen
  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.imgWrapper}>
        <Image
          style={styles.image}
          source={{
            uri:
              "https://www.heikaus-asset.com/wp-content/uploads/2018/01/Test_Starbucks-1024x731.jpg"
          }}
        ></Image>
      </View>
      <View style={styles.contentWrapper}>
        <Text style={styles.categoryTag}>{data.Desc}</Text>
        <Text style={styles.headerText}>{data.Name}</Text>
        <RatingStars colorFill={color.brand.normal} rating={data.Rating} />
        <FeatureIcons features={data.Features} />
        <View style={styles.horizontalRule} />
        <Text style={styles.secondaryHeaderText}>Kommentare</Text>
      </View>
      <FlatList
        keyExtractor={item => item.ID}
        data={socketReviews}
        style={styles.contentWrapper}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => <ReviewCard reviewData={item} />}
      />
    </SafeAreaView>
  );
};

// Define the styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white
  },
  imgWrapper: {
    alignItems: "center",
    marginBottom: 20
  },
  image: {
    height: 275,
    width: 500
  },
  contentWrapper: {
    paddingHorizontal: 30,
    marginBottom: 25
  },
  categoryTag: {
    fontSize: 20,
    fontWeight: "300",
    textTransform: "uppercase"
  },
  headerText: {
    fontSize: 30,
    fontWeight: "700"
  },
  horizontalRule: {
    borderBottomColor: color.gray.dark,
    borderBottomWidth: 1,
    marginVertical: 20
  },
  secondaryHeaderText: {
    fontSize: 20,
    fontWeight: "500"
  }
});

// Export the above defined screen
export default LocationDetailsScreen;
