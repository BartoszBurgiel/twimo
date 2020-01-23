// Import core functionality dependencies
import React from "react";
import { View, ScrollView, Text, StyleSheet, Image } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Import custom built components
import RatingStars from "../components/RatingStars";
import SupplyIcons from "../components/SupplyIcons";
import ReviewCard from "../components/ReviewCard";

const dummyReviews = [
  {
    name: "Bartosz Burgiel",
    avatar: "https://api.adorable.io/avatars/50/bartosz@mail.com.png",
    rating: 5,
    content:
      "Klasse Handwerksbäcker, gemütliches und stilvolles Café mit hervorragenden Backwaren und bestem Kaffee. Unbedingt den Cafe Latte probieren - ein echter Genuss an diesen kalten Winterertagen!"
  },
  {
    name: "Louis Beul",
    avatar: "https://api.adorable.io/avatars/50/louis@mail.com.png",
    rating: 4,
    content:
      "Klasse Cafe, gemütliches und stilvolles Ambiente mit hervorragenden Backwaren und gutem Kaffee. Unbedingt den Cafe Latte probieren - ein echter Genuss an diesen kalten Winterertagen!"
  }
];

// Define Component
const LocationDetailsScreen = () => {
  return (
    <View style={styles.container}>
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
        <Text style={styles.categoryTag}>coffeeshop</Text>
        <Text style={styles.headerText}>Starbucks Montabaur</Text>
        <RatingStars colorFill={color.brand.normal} rating={4} />
        <SupplyIcons />
        <View style={styles.horizontalRule} />
        <Text style={styles.secondaryHeaderText}>Reviews</Text>
      </View>
      <ScrollView
        style={styles.contentWrapper}
        showsVerticalScrollIndicator={false}
      >
        <ReviewCard reviewData={dummyReviews[0]} />
        <ReviewCard reviewData={dummyReviews[1]} />
      </ScrollView>
    </View>
  );
};

// Define styles
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

// Export the above defined component
export default LocationDetailsScreen;
