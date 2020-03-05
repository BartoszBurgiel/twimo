// Import core functionality dependencies from react & react-native
import React, { useState } from "react";
import {
  View,
  ScrollView,
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
  //FIXME: Virtualized Lists Warning

  // Set up component's state
  const [reviews, setReviews] = useState(0);

  // Grab locationData from the props and store it in data const
  const data = props.navigation.state.params.locationData;

  // Fetch review data from API
  fetch("http://127.0.0.1:5500/frontend/data/reviews.json")
    .then(response => response.json())
    .then(data => setReviews(data));

  return (
    <ScrollView style={styles.container}>
      <View style={styles.imgWrapper}>
        <Image
          style={styles.image}
          source={{
            uri: data.picture
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
        data={reviews}
        style={styles.contentWrapper}
        showsVerticalScrollIndicator={false}
        renderItem={({ item }) => <ReviewCard reviewData={item} />}
      />
    </ScrollView>
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
